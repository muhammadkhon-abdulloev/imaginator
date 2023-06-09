package handler

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	v1 "github.com/muhammadkhon-abdulloev/imaginator/gen/go/imaginator/v1"
	"github.com/muhammadkhon-abdulloev/imaginator/pkg/storage"
	"github.com/muhammadkhon-abdulloev/pkg/logger"
	"io"
)

type Handler struct {
	duLimitCh  chan struct{} // download/upload limit chan
	lafLimitCh chan struct{} // list all files limit chan
	_storage   *storage.Storage
	lg         *logger.Logger
	v1.UnimplementedImaginatorServer
}

func NewHandler(
	duLimit,
	lafLimit int,
	_storage *storage.Storage,
	lg *logger.Logger,
) *Handler {
	return &Handler{
		duLimitCh:  make(chan struct{}, duLimit),
		lafLimitCh: make(chan struct{}, lafLimit),
		_storage:   _storage,
		lg:         lg,
	}
}

func (h *Handler) UploadFile(_ context.Context, req *v1.UploadFileRequest) (*v1.UploadFileResponse, error) {
	h.duLimitCh <- struct{}{}
	defer func() {
		<-h.duLimitCh
	}()

	file, err := h._storage.Upload(req.GetFilename(), req.GetData())
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
		Data: file,
	}, nil
}

func (h *Handler) ListAllFiles(_ context.Context, _ *v1.ListAllFilesMessage) (*v1.ListAllFilesMessage, error) {
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

	for _, file := range files {
		resFiles = append(resFiles, &v1.Files{
			Filename:  file.Name,
			CreatedAt: file.CreatedAt.UnixMilli(),
			UpdatedAt: file.LastUpdate.UnixMilli(),
		})
	}

	return &v1.ListAllFilesMessage{Files: resFiles}, nil
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

	data, err := io.ReadAll(bf)
	if err != nil {
		err = fmt.Errorf("io.ReadAll: %w", err)
		h.lg.Error(err)
		return err
	}

	file, err := h._storage.Upload(filename, data)
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

	file, err := h._storage.Download(req.GetFilename())
	if err != nil {
		err = fmt.Errorf("h._storage.Download: %w", err)
		h.lg.Error(err)
		return err
	}

	reader := bufio.NewReader(bytes.NewReader(file))
	for {
		chunk := make([]byte, req.GetChunkSize())
		_, err = reader.Read(chunk)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				err = fmt.Errorf("reader.Read: %w", err)
				h.lg.Error(err)
				return err
			}
			break
		}
		err = conn.Send(&v1.DownloadFileResponse{
			Data: chunk,
		})
		if err != nil {
			err = fmt.Errorf("conn.Send: %w", err)
			h.lg.Error(err)
			return err
		}
	}

	return nil
}