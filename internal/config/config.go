package config

import (
	"context"
	"fmt"
	"github.com/muhammadkhon-abdulloev/pkg/logger"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/fx"
)

type Config struct {
	Logger logger.Config
	Server Server
}

var (
	_c     = &Config{}
	Option = fx.Provide(context.Background, Init)
)

func GetConfig() *Config {
	return _c
}

type Server struct {
	Port                string `env:"PORT,default=5555"`
	ImagesPath          string `env:"IMAGES_PATH,default=assets"`
	UploadDownloadLimit int    `env:"UPLOAD_DOWNLOAD_LIMIT,default=10"`
	ListAllFilesLimit   int    `env:"LIST_ALL_FILES_LIMIT,default=100"`
}

func Init(ctx context.Context) (*Config, error) {
	if err := envconfig.Process(ctx, _c); err != nil {
		return nil, fmt.Errorf("envconfig.Process: %w", err)
	}

	return _c, nil
}
