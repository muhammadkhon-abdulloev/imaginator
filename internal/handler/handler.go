package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	v1 "github.com/muhammadkhon-abdulloev/imaginator/gen/go/imaginator/v1"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/config"
	"github.com/muhammadkhon-abdulloev/imaginator/pkg/storage"
	"github.com/muhammadkhon-abdulloev/imaginator/pkg/storage/disk"
	"github.com/muhammadkhon-abdulloev/pkg/logger"
	"go.uber.org/fx"
	"io"
)

type Handler struct {
	duLimitCh  chan struct{} // download/upload limit chan
	lafLimitCh chan struct{} // list all files limit chan
	_storage   storage.StorageManager
	lg         *logger.Logger
	v1.UnimplementedImaginatorServer
}

var (
	Option = fx.Provide(New)
)

type Params struct {
	fx.In
	Config  *config.Config
	Logger  *logger.Logger
	Storage *disk.Storage
}

func New(p Params) *Handler {
	return &Handler{
		duLimitCh:  make(chan struct{}, p.Config.Server.UploadDownloadLimit),
		lafLimitCh: make(chan struct{}, p.Config.Server.ListAllFilesLimit),
		_storage:   p.Storage,
		lg:         p.Logger,
	}
}

func (h *Handler) UploadFile(_ context.Context, req *v1.UploadFileRequest) (*v1.UploadFileResponse, error) {
	h.duLimitCh <- struct{}{}
	defer func() {
		<-h.duLimitCh
	}()
	filename := req.GetFilename()

	if len(filename) < 3 {
		return nil, ErrInvalidFileName
	}

	file, err := h._storage.Upload(filename, req.GetData())
	if err != nil {
		err = fmt.Errorf("h._storage.Upload: %w", err)
		h.lg.Error(err)
		return nil, err
	}

	return &v1.UploadFileResponse{
		Filename: file.Name,
		Filesize: file.Size,
	}, nil
}

func (h *Handler) DownloadFile(_ context.Context, req *v1.DownloadFileRequest) (*v1.DownloadFileResponse, error) {
	h.duLimitCh <- struct{}{}
	defer func() {
		<-h.duLimitCh
	}()

	file, err := h._storage.Download(req.GetFilename())
	if err != nil {
		err = fmt.Errorf("h._storage.Download: %w", err)
		h.lg.Error(err)

		return nil, err
	}

	return &v1.DownloadFileResponse{
		Data:     file.Bytes,
		Checksum: file.CheckSum,
	}, nil
}

func (h *Handler) ListAllFiles(_ context.Context, req *v1.ListAllFilesRequest) (*v1.ListAllFilesResponse, error) {
	var resFiles []*v1.Files
	h.lafLimitCh <- struct{}{}
	defer func() {
		<-h.lafLimitCh
	}()

	files, err := h._storage.ListAllFiles()
	if err != nil {
		err = fmt.Errorf("h._storage.ListAllFiles: %w", err)
		h.lg.Error(err)
		return nil, err
	}

	for _, file := range h._storage.Paginate(req.GetLimit(), req.GetOffset(), files) {
		resFiles = append(resFiles, &v1.Files{
			Filename:  file.Name,
			CreatedAt: file.CreatedAt.UnixMilli(),
			UpdatedAt: file.LastUpdate.UnixMilli(),
		})
	}

	return &v1.ListAllFilesResponse{
		Files:    resFiles,
		TotalLen: int64(len(files)),
	}, nil
}

func (h *Handler) UploadFileByChunk(conn v1.Imaginator_UploadFileByChunkServer) error {
	h.duLimitCh <- struct{}{}
	defer func() {
		<-h.duLimitCh
	}()
	var (
		filename = ""
		bf       = &bytes.Buffer{}
	)

	for {
		req, err := conn.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				err = fmt.Errorf("conn.Recv:%w", err)
				h.lg.Error(err)
				return err
			}
			break
		}
		bf.Write(req.GetData())
		filename = req.GetFilename()
	}

	if len(filename) < 3 {
		return ErrInvalidFileName
	}

	file, err := h._storage.Upload(filename, bf.Bytes())
	if err != nil {
		err = fmt.Errorf("h._storage.Upload: %w", err)
		h.lg.Error(err)
		return err
	}

	err = conn.SendAndClose(&v1.UploadFileResponse{
		Filename: file.Name,
		Filesize: file.Size,
	})

	if err != nil {
		err = fmt.Errorf("conn.SendAndClose: %w", err)
		h.lg.Error(err)
		return err
	}
	return nil
}

func (h *Handler) DownloadFileByChunk(req *v1.DownloadFileRequest, conn v1.Imaginator_DownloadFileByChunkServer) error {
	h.duLimitCh <- struct{}{}
	defer func() {
		<-h.duLimitCh
	}()
	err := h.streamFileAsync(req.GetFilename(), int(req.GetChunkBuffSize()), conn)
	if err != nil {
		err = fmt.Errorf("%w", err)
		h.lg.Error(err)
		return err
	}

	return nil
}
