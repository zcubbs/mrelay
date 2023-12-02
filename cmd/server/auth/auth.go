package auth

import (
	"crypto/subtle"
	"fmt"
	"github.com/zcubbs/mrelay/cmd/server/db"
	"net/http"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement auth middleware
		next.ServeHTTP(w, r)
	})
}

// BasicAuth implements a simple middleware handler for adding basic http auth to a route.
func BasicAuth(realm string, store db.AuthStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok {
				basicAuthFailed(w, realm)
				return
			}

			dbUser, err := store.GetUserByUsername(user)
			if err != nil || subtle.ConstantTimeCompare([]byte(pass), []byte(dbUser.Password)) != 1 {
				basicAuthFailed(w, realm)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func basicAuthFailed(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}
