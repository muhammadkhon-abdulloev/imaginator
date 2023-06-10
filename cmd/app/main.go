package main

import (
	"context"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/config"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/handler"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/middleware"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/server"
	"github.com/muhammadkhon-abdulloev/imaginator/pkg/logger"
	"github.com/muhammadkhon-abdulloev/imaginator/pkg/storage/disk"
	"go.uber.org/fx"
)

func main() {
	mainModules := fx.Options(
		fx.Provide(context.Background),
		config.Option,
		logger.Option,
		disk.Option,
		handler.Option,
		middleware.Option,
		server.Option,
	)

	fx.New(mainModules).Run()
}
