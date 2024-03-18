package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/akrovv/filmlibrary/internal/adapters/postgresqldb"
	"github.com/akrovv/filmlibrary/internal/adapters/redisdb"
	"github.com/akrovv/filmlibrary/internal/config"
	"github.com/akrovv/filmlibrary/internal/controllers/restapi"
	"github.com/akrovv/filmlibrary/internal/controllers/restapi/middleware"
	"github.com/akrovv/filmlibrary/internal/service"
	"github.com/akrovv/filmlibrary/pkg/hasher"
	"github.com/akrovv/filmlibrary/pkg/logger"
	"github.com/casbin/casbin/v2"
	md "github.com/go-openapi/runtime/middleware"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

const (
	path     = "."
	filename = ".env"
	model    = "./rbac_model.conf"
	policy   = "./rbac_policy.csv"
)

// @title FilmLibrary
// @version 1.0
// @description API server for FilmLibrary

// @host      localhost:8080
// @BasePath  /

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
		return
	}

	enforcer, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		logger.Info(err)
		return
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Info(err)
		return
	}

	cfg, err := config.NewConfig(path, filename)
	if err != nil {
		logger.Info(err)
		return
	}

	ctxRedis := context.Background()
	dsnRedis := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr: dsnRedis,
		DB:   0,
	})

	if err = client.Ping(ctxRedis).Err(); err != nil {
		logger.Info(err)
		return
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort, cfg.DBUser, cfg.DBPassword,
		cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Info(err)
		return
	}

	db.SetMaxOpenConns(10)

	if err = db.Ping(); err != nil {
		logger.Info(err)
		return
	}

	var (
		userHasher    = hasher.NewHasher([]byte("secret"))
		sessionHasher = hasher.NewHasher([]byte("session"))
	)

	var (
		actorStorage   = postgresqldb.NewActorStorage(db)
		movieStorage   = postgresqldb.NewMovieStorage(db)
		userStorage    = postgresqldb.NewUserStorage(db, userHasher)
		sessionStorage = redisdb.NewSessionStorage(ctxRedis, client, sessionHasher)
	)

	var (
		actorService   = service.NewActorService(actorStorage)
		movieService   = service.NewMovieService(movieStorage)
		userService    = service.NewUserService(userStorage)
		sessionService = service.NewSessionService(sessionStorage)
	)

	var (
		actorController = restapi.NewActorController(logger, actorService)
		movieController = restapi.NewMovieController(logger, movieService)
		userController  = restapi.NewUserController(logger, userService, sessionService)
	)

	mux := http.NewServeMux()

	opts := md.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := md.Redoc(opts, nil)

	mux.Handle("/docs", sh)
	mux.Handle("/swagger.yaml", http.FileServer(http.Dir("./docs/")))

	mux.HandleFunc("/register", userController.Register)
	mux.HandleFunc("/login", userController.Login)

	mux.HandleFunc("/actor", actorController.ManagePath)

	mux.HandleFunc("/movie/all", movieController.GetOrderedList)
	mux.HandleFunc("/movie", movieController.ManagePath)

	var (
		roleMiddleware   = middleware.Role(mux, enforcer)
		authMiddleware   = middleware.Auth(roleMiddleware, sessionService)
		loggerMiddleware = middleware.Logger(authMiddleware, logger)
	)

	logger.Infof("starting on :%s", cfg.ServerPort)
	if err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServerPort), loggerMiddleware); err != nil {
		logger.Info(err)
		return
	}
}
