package server

import (
	"context"
	"fmt"
	"github.com/muhammadkhon-abdulloev/imaginator/gen/go/imaginator/v1"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/config"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/handler"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/middleware"
	"github.com/muhammadkhon-abdulloev/pkg/logger"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	lg               *logger.Logger
	mdw              *middleware.Middleware
	imaginatorServer v1.ImaginatorServer
}

var Option = fx.Invoke(New)

type Params struct {
	fx.In
	Logger           *logger.Logger
	Middleware       *middleware.Middleware
	ImaginatorServer *handler.Handler
}

func New(
	lc fx.Lifecycle,
	p Params,
) {
	s := &Server{
		lg:               p.Logger,
		mdw:              p.Middleware,
		imaginatorServer: p.ImaginatorServer,
	}
	var srv *grpc.Server
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			srv, err = s.Run()
			if err != nil {
				return err
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.GracefulStop()
			return nil
		},
	})
}

func (s Server) Run() (*grpc.Server, error) {
	l, err := net.Listen("tcp", ":"+config.GetConfig().Server.Port)
	if err != nil {
		return nil, fmt.Errorf("net.Listen: %w", err)
	}
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			s.mdw.UnaryLoggingInterceptor(),
			s.mdw.UnaryRecoverInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			s.mdw.StreamLoggingInterceptor(),
			s.mdw.StreamRecoverInterceptor(),
		),
	)

	v1.RegisterImaginatorServer(srv, s.imaginatorServer)

	go func() {
		if err := srv.Serve(l); err != nil {
			s.lg.Error(fmt.Errorf("srv.Serve: %w", err).Error())
		}
	}()

	return srv, nil
}
