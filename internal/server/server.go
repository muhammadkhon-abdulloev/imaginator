package server

import (
	"fmt"
	"github.com/muhammadkhon-abdulloev/imaginator/gen/go/imaginator/v1"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/config"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/handler"
	"github.com/muhammadkhon-abdulloev/imaginator/pkg/storage/disk"
	"github.com/muhammadkhon-abdulloev/pkg/logger"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	lg *logger.Logger
}

func NewServer(lg *logger.Logger) *Server {
	return &Server{
		lg: lg,
	}
}

func (s Server) Run() error {
	l, err := net.Listen("tcp", ":"+config.GetConfig().Server.Port)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	st := disk.NewStorage(config.GetConfig().Server.ImagesPath)
	h := handler.NewHandler(
		config.GetConfig().Server.UploadDownloadLimit,
		config.GetConfig().Server.ListAllFilesLimit,
		st,
		s.lg,
	)

	srv := grpc.NewServer()
	v1.RegisterImaginatorServer(srv, h)

	if err = srv.Serve(l); err != nil {
		return fmt.Errorf("srv.Serve: %w", err)
	}

	return nil
}
