package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"robocar-webserver/src/packages/appConfig"
	"robocar-webserver/src/packages/basicAuth"
	"robocar-webserver/src/packages/ngrokTunnel"
	"robocar-webserver/src/packages/proxy"
	"robocar-webserver/src/packages/streaming"
)

//go:embed all:static
var staticAssets embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(staticAssets, "static")
}

func main() {
	var err error
	_config := appConfig.Init()

	ctlProxy, _ := proxy.NewProxy(_config.RobocarCtrUrl)

	assets, _ := Assets()
	fileServer := http.FileServer(http.FS(assets))

	proxy.Register("/ctl/", ctlProxy)

	streamHandler := streaming.MakeMultiplexedStreamHandler()
	http.HandleFunc("/stream", streamHandler)

	http.HandleFunc("/", basicAuth.Middleware(fileServer.ServeHTTP))

	if _config.NgrokEnabled {
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
