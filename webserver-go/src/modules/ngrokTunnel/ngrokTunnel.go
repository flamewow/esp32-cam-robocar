package ngrokTunnel

import (
	"context"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"os"
)

func CreateTunnel(ctx context.Context) (ngrok.Tunnel, error) {
	return ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithDomain(os.Getenv("NGROK_DOMAIN")),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
}
