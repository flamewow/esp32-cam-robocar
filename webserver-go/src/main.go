package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	targetUrl, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(targetUrl), nil
}

func HandleReverseProxy(prefix string, proxy *httputil.ReverseProxy) {
	http.Handle(prefix, http.StripPrefix(prefix, proxy))
}

//go:embed all:static
var staticAssets embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(staticAssets, "static")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	robocarIp := os.Getenv("ROBOCAR_IP")

	//// Basic authentication middleware
	//auth := func(handler http.Handler) http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		user, pass, ok := r.BasicAuth()
	//		if !ok || user != "admin" || pass != "robot" {
	//			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	//			w.WriteHeader(http.StatusUnauthorized)
	//			return
	//		}
	//		handler.ServeHTTP(w, r)
	//	})
	//}

	streamProxy, _ := NewProxy(fmt.Sprintf("http://%s:81", robocarIp))
	ctlProxy, _ := NewProxy(fmt.Sprintf("http://%s:80", robocarIp))

	assets, _ := Assets()
	fileServer := http.FileServer(http.FS(assets))

	HandleReverseProxy("/ctl/", ctlProxy)
	HandleReverseProxy("/stream/", streamProxy)

	http.Handle("/", fileServer)

	fmt.Println("Started listening at http://0.0.0.0:4000")
	err = http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
