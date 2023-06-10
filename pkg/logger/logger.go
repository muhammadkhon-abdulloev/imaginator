package logger

import (
	"fmt"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/config"
	"github.com/muhammadkhon-abdulloev/pkg/logger"
	"go.uber.org/fx"
)

var Option = fx.Provide(New)

type Params struct {
	fx.In
	Config *config.Config
}

func New(params Params) (*logger.Logger, error) {
	lg := logger.NewLogger()
	if err := lg.InitLogger(params.Config.Logger); err != nil {
		return nil, fmt.Errorf("lg.InitLogger: %w", err)
	}

	return lg, nil
}
