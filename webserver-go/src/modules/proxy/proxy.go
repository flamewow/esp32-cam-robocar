package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"robocar-webserver/src/modules/basicAuth"
)

func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	targetUrl, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(targetUrl), nil
}

func HandleReverseProxy(prefix string, proxy *httputil.ReverseProxy) {
	http.HandleFunc(prefix, basicAuth.Middleware(http.StripPrefix(prefix, proxy).ServeHTTP))
}
