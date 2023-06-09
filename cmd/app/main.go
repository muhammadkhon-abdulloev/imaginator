package main

import (
	"context"
	"fmt"
	"github.com/muhammadkhon-abdulloev/imaginator/config"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/server"
	"github.com/muhammadkhon-abdulloev/pkg/logger"
)

func main() {
	cfg, err := config.Init(context.Background())
	if err != nil {
		panic(fmt.Errorf("config.Init: %w", err).Error())
	}

	lg := logger.NewLogger()
	if err = lg.InitLogger(cfg.Logger); err != nil {
		panic(fmt.Errorf("lg.InitLogger: %w", err).Error())
	}

	srv := server.NewServer(lg)
	if err = srv.Run(); err != nil {
		panic(fmt.Errorf("srv.Run: %w", err).Error())
	}
}
