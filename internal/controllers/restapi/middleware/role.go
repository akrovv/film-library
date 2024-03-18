package middleware

import (
	"net/http"

	"github.com/akrovv/filmlibrary/internal/domain"
	"github.com/casbin/casbin/v2"
)

func Role(next http.Handler, enforcer *casbin.Enforcer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userContext domain.UserContext = "user"

		ctxUser := r.Context().Value(userContext)
		sub := "anonymous"

		if ctxUser != nil {
			user := ctxUser.(*domain.User)
			sub = "user"

			if user.IsAdmin {
				sub = "admin"
			}
		}

		obj := r.URL.Path
		act := r.Method
		ok, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

/*
user
curl -v localhost:8080/docs
curl -v -X POST -H "Content-Type: application/json" -d '{"username": "akro", "password": "akro"}' localhost:8080/register
curl -v --cookie "session-id=73657373696f6e18890d6a9ce0cdbb5a8ce0c14a795525" -X GET -H "Content-Type: application/json" -d '{"username": "akro", "password": "akro"}' localhost:8080/actor
curl -v --cookie "session-id=73657373696f6e18890d6a9ce0cdbb5a8ce0c14a795525" -X POST -H "Content-Type: application/json" -d '{"username": "akro", "password": "akro"}' localhost:8080/actor
admin
movie:
73657373696f6e18890d6a9ce0cdbb5a8ce0c14a795525
curl -v --cookie -X POST -H "Content-Type: application/json" -d '{"username": "admin", "password": "admin"}' localhost:8080/login
curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X POST -H "Content-Type: application/json" -d '{"movie_title": "Avatar", "description": "good film!", "release_date": "2007-02-02T00:00:00Z", "rating": 9, "actors": [1, 2]}' localhost:8080/movie
curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X PUT -H "Content-Type: application/json" -d '{"movie_id": 2, "movie_title": "Avatar", "description": "good film!", "release_date": "2007-02-02T00:00:00Z", "rating": 6, "actors": [1]}' localhost:8080/movie
curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X GET "localhost:8080/movie?title=ava"
curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X GET "localhost:8080/movie/all"

curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X GET localhost:8080/actor
curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X POST -H "Content-Type: application/json" -d '{"actor_name": "Artyom Blokhin", "gender": "Male", "date_of_birth": "2007-02-02T00:00:00Z"}' localhost:8080/actor
curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X PUT -H "Content-Type: application/json" -d '{"actor_id": 4, "actor_name": "Blokha", "gender": "Male", "date_of_birth": "2007-02-02T00:00:00Z"}' localhost:8080/actor
curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X DELETE -H "Content-Type: application/json" -d '{"actor_id": 4}' localhost:8080/actor
curl -v -X POST -H "Content-Type: application/json" -d '{"username": "admin", "password": "admin"}' localhost:8080/login

curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X POST -H "Content-Type: application/json" -d '{"actor_name": "Mister", "gender": "Male", "date_of_birth": "2007-02-02T00:00:00Z"}' localhost:8080/actor
curl -v --cookie "session-id=73657373696f6e21232f297a57a5a743894a0e4a801fc3" -X POST -H "Content-Type: application/json" -d '{"movie_title": "Skazka", "description": "good film!", "release_date": "2007-02-02T00:00:00Z", "rating": 10, "actors": [1]}' localhost:8080/movie
*/
