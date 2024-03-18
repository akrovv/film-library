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
