package restapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/akrovv/filmlibrary/pkg/logger"
	"github.com/akrovv/filmlibrary/pkg/sender"
)

type userController struct {
	logger         logger.Logger
	userService    UserService
	sessionService SessionService
}

func NewUserController(logger logger.Logger,
	userService UserService,
	sessionService SessionService) *userController {
	return &userController{
		logger:         logger,
		userService:    userService,
		sessionService: sessionService,
	}
}

func userUnmarshal(data []byte) (*domain.CRUser, error) {
	user := &domain.CRUser{}

	err := json.Unmarshal(data, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// @Summary Register
// @Description Register a new user
// @Tags		 user
// @Accept       json
// @Produce      json
// @Param request body domain.CRUser true "request"
// @Success 201 {object} sender.JSONResponse
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /register [post]
func (c *userController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		c.logger.Infof("request didnt contain application/json")
		sender.ErrorJSON(w, domain.ErrRequest, http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		c.logger.Infof("io.ReadAll error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	crUserDTO, err := userUnmarshal(data)
	if err != nil {
		c.logger.Infof("json.Unmarshal error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = c.userService.Register(crUserDTO)
	if err != nil {
		c.logger.Infof("c.UserService.Register %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	createSessionDTO := domain.CreateSession{
		Username: crUserDTO.Username,
	}

	id, err := c.sessionService.Create(&createSessionDTO)
	if err != nil {
		c.logger.Infof("c.SessionService.Create %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	c.logger.Infof("created session for user: [%s] with session-id: [%s]", crUserDTO.Username, id)
	http.SetCookie(w, &http.Cookie{
		Name:     "session-id",
		Value:    id,
		Expires:  time.Now().Add(time.Hour * 8),
		HttpOnly: true,
	})

	if err = sender.WriteJSON(w, http.StatusCreated, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "OK",
		Message: "user was created",
	}); err != nil {
		c.logger.Infof("WriteJSON %w", err)
		return
	}
}

// @Summary Login
// @Description Log in with user credentials
// @Tags		 user
// @Accept       json
// @Produce      json
// @Param request body domain.CRUser true "request"
// @Success 201 {object} sender.JSONResponse
// @Failure 400 {object} sender.JSONResponse
// @Failure 500 {object} sender.JSONResponse
// @Router       /login [post]
func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		c.logger.Infof("request didnt contain application/json")
		sender.ErrorJSON(w, domain.ErrRequest, http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		c.logger.Infof("io.ReadAll error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	crUserDTO, err := userUnmarshal(data)
	if err != nil {
		c.logger.Infof("json.Unmarshal error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user, err := c.userService.Login(crUserDTO)
	if err != nil {
		c.logger.Infof("c.UserService.Login error: %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	createSessionDTO := domain.CreateSession{
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
	}

	id, err := c.sessionService.Create(&createSessionDTO)
	if err != nil {
		c.logger.Infof("c.SessionService.Create %w", err)
		sender.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	c.logger.Infof("created session for user: [%s] with session-id: [%s]", user.Username, id)

	http.SetCookie(w, &http.Cookie{
		Name:     "session-id",
		Value:    id,
		Expires:  time.Now().Add(time.Hour * 8),
		HttpOnly: true,
	})

	if err = sender.WriteJSON(w, http.StatusCreated, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "OK",
		Message: "user was login",
	}); err != nil {
		c.logger.Infof("WriteJSON %w", err)
		return
	}
}
