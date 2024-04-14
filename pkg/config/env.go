package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type Environment struct {
	Port       int64 `env:"PORT,default=8080"`
	Production bool  `env:"PRODUCTION,default=false"`

	LocalEndpoint    string `env:"LOCAL_ENDPOINT"`
	LocalAccessKeyID string `env:"LOCAL_ACCESS_KEY_ID"`
	LocalSecretKey   string `env:"LOCAL_SECRET_KEY"`
	LocalUseSSL      bool   `env:"LOCAL_USE_SSL,default=false"`

	RemoteEndpoint    string `env:"REMOTE_ENDPOINT"`
	RemoteAccessKeyID string `env:"REMOTE_ACCESS_KEY_ID"`
	RemoteSecretKey   string `env:"REMOTE_SECRET_KEY"`
	RemoteUseSSL      bool   `env:"REMOTE_USE_SSL,default=true"`

	Extras env.EnvSet
}

func GetConfig() (*Environment, error) {
	var environment Environment

	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	environment.Extras = es

	return &environment, nil
}
