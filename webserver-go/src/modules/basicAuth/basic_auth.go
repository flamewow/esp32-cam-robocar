package basicAuth

import (
	"net/http"
	"os"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authEnabled := os.Getenv("AUTH_ENABLED")

		if authEnabled != "TRUE" {
			next.ServeHTTP(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || user != os.Getenv("AUTH_USER") || pass != os.Getenv("AUTH_PASSWORD") {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
