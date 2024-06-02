package appConfig

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	RobocarIp      string `env:"ROBOCAR_IP" envDefault:"localhost"`
	NgrokEnabled   bool   `env:"NGROK_ENABLED" envDefault:"false"`
	NgrokDomain    string `env:"NGROK_DOMAIN"`
	NgrokAuthToken string `env:"NGROK_AUTH_TOKEN"`

	AuthEnabled  bool   `env:"AUTH_ENABLED" envDefault:"false"`
	AuthUser     string `env:"AUTH_USER"`
	AuthPassword string `env:"AUTH_PASSWORD"`
}

var AppConfig Config

// Init AppConfig structure (load from env variables, .env file etc), call in in main()
func Init() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	AppConfig, err := env.ParseAs[Config]()

	if err != nil {
		panic(fmt.Errorf("failed to parse appConfig: %v", err))
	}

	return &AppConfig
}

func Load() *Config {
	return &AppConfig
}
