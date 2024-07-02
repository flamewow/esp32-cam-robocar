package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"robocar-webserver/src/packages/basicAuth"
)

func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	targetUrl, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(targetUrl), nil
}

func removeCookieHeader(original func(*http.Request)) func(*http.Request) {
	return func(req *http.Request) {
		original(req)
		req.Header.Del("Cookie")
	}
}

func Register(prefix string, proxy *httputil.ReverseProxy) {
	proxy.Director = removeCookieHeader(proxy.Director)

	http.HandleFunc(prefix, basicAuth.Middleware(http.StripPrefix(prefix, proxy).ServeHTTP))
}
