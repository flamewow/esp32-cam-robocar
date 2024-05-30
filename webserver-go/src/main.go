package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"robocar-webserver/src/modules/basicAuth"
	"robocar-webserver/src/modules/ngrokTunnel"

	"github.com/joho/godotenv"
	"robocar-webserver/src/modules/proxy"
)

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

	streamProxy, _ := proxy.NewProxy(fmt.Sprintf("http://%s:81", robocarIp))
	ctlProxy, _ := proxy.NewProxy(fmt.Sprintf("http://%s:80", robocarIp))

	assets, _ := Assets()
	fileServer := http.FileServer(http.FS(assets))

	proxy.HandleReverseProxy("/ctl/", ctlProxy)
	proxy.HandleReverseProxy("/stream/", streamProxy)

	http.HandleFunc("/", basicAuth.Middleware(fileServer.ServeHTTP))

	if os.Getenv("NGROK_ENABLED") == "true" {
		ngrokListener, err := ngrokTunnel.CreateTunnel(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Ingress established at:", ngrokListener.URL())
		err = http.Serve(ngrokListener, nil)
	} else {
		log.Println("Ingress established at: http://localhost:4000")
		err = http.ListenAndServe(":4000", nil)
	}

	if err != nil {
		log.Fatal(err)
	}
}
