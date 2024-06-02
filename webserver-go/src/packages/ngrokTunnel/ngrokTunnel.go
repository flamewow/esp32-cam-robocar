package ngrokTunnel

import (
	"context"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"robocar-webserver/src/packages/appConfig"
)

var _config = appConfig.Load()

func CreateTunnel(ctx context.Context) (ngrok.Tunnel, error) {
	return ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithDomain(_config.NgrokDomain),
		),
		ngrok.WithAuthtoken(_config.NgrokAuthToken),
	)
}
