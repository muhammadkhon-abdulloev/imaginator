package config

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Server Server
}

var _c = &Config{}

func GetConfig() *Config {
	return _c
}

type Server struct {
	Port string `env:"PORT,default=5555"`
}

func Init(ctx context.Context) (*Config, error) {
	if err := envconfig.Process(ctx, _c); err != nil {
		return nil, fmt.Errorf("envconfig.Process: %w", err)
	}

	return _c, nil
}
