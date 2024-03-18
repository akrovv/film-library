package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/akrovv/filmlibrary/internal/controllers/restapi"
	"github.com/akrovv/filmlibrary/internal/domain"
)

func Auth(next http.Handler, sessionService restapi.SessionService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session-id")
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			return
		}

		path := r.URL.Path
		if err != nil {
			if path == "/register" || path == "/login" || path == "/swagger.yaml" || path == "/docs" {
				next.ServeHTTP(w, r)
				return
			}

			http.Redirect(w, r, "/register", http.StatusPermanentRedirect)
			return
		}

		value := cookie.Value
		getSessionDTO := domain.GetSession{
			Username: value,
		}

		user, err := sessionService.Get(&getSessionDTO)
		if err != nil {
			return
		}

		var userContext domain.UserContext = "user"
		ctx := context.WithValue(r.Context(), userContext, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
