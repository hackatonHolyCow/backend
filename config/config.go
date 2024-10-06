package config

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Application Application
		Databases   Databases
		MercadoPago MercadoPago
		OpenAI      OpenAI
	}

	Application struct {
		Environment string `env:"ENVIRONMENT,required=true"`
		Port        int32  `env:"PORT,required=true"`
	}

	Databases struct {
		PostgresDSN string `env:"POSTGRES_DSN,required=true"`
	}

	MercadoPago struct {
		Token string `env:"MP_ACCESS_TOKEN,required=true"`
	}

	OpenAI struct {
		APIKey string `env:"OPENAI_API_KEY,required=true"`
	}
)

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("not .env found")
	}

	var conf Config
	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
