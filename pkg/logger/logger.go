package logger

import (
	"fmt"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/config"
	"github.com/muhammadkhon-abdulloev/pkg/logger"
	"go.uber.org/fx"
)

var Option = fx.Provide(New)

func New(cfg *config.Config) (*logger.Logger, error) {
	lg := logger.NewLogger()
	if err := lg.InitLogger(cfg.Logger); err != nil {
		return nil, fmt.Errorf("lg.InitLogger: %w", err)
	}

	return lg, nil
}
