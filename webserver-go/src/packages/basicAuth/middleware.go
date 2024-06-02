package basicAuth

import (
	"net/http"
	"robocar-webserver/src/packages/appConfig"
)

var _config = appConfig.Load()

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authEnabled := _config.AuthEnabled

		if !authEnabled {
			next.ServeHTTP(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || user != _config.AuthUser || pass != _config.AuthPassword {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
