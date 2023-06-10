package handler

import (
	"bufio"
	"errors"
	"fmt"
	v1 "github.com/muhammadkhon-abdulloev/imaginator/gen/go/imaginator/v1"
	"golang.org/x/sync/errgroup"
	"io"
	"os"
)

const maxActiveGoroutines = 10

type chunk struct {
	bufSize int
	offset  int
}

// streamFile - streams file synchronously
func (h *Handler) streamFile(
	filename string,
	chunkBuffSize int,
	conn v1.Imaginator_DownloadFileByChunkServer,
) error {
	file, err := h._storage.GetFile(filename)
	if err != nil {
		err = fmt.Errorf("h._storage.Download: %w", err)
		h.lg.Error(err)
		return err
	}

	reader := bufio.NewReader(file)

	defer file.Close()
	for i := 0; ; i++ {
		buf := make([]byte, chunkBuffSize)
		n, err := reader.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return fmt.Errorf("file.Read: %w", err)
		}
		err = conn.Send(&v1.DownloadFileResponse{
			Data:   buf,
			Offset: int64(i),
		})
		if err != nil {
			return fmt.Errorf("conn.Send: %w", err)
		}
		buf = buf[:n]
	}

	return nil
}

// streamFileAsync - streams file asynchronously using goroutines and error group
func (h *Handler) streamFileAsync(
	filename string,
	chunkBuffSize int,
	conn v1.Imaginator_DownloadFileByChunkServer,
) error {
	file, err := h._storage.GetFile(filename)
	if err != nil {
		err = fmt.Errorf("h._storage.Download: %w", err)
		h.lg.Error(err)
		return err
	}

	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("file.Stat: %w", err)
	}

	filesize := int(fileStat.Size())

	// quantity of chunks
	chunksQ := filesize / chunkBuffSize
	if filesize%chunkBuffSize != 0 {
		chunksQ++
	}

	chunks := make([]chunk, chunksQ)

	for i := 0; i < chunksQ; i++ {
		chunks[i].offset = i * chunkBuffSize
		chunks[i].bufSize = chunkBuffSize
	}

	eg := &errgroup.Group{}
	eg.SetLimit(maxActiveGoroutines)
	for i := 0; i < chunksQ; i++ {
		i := i
		eg.Go(func() error {
			return h.sendToStream(i, chunks, file, conn)
		})
	}

	return eg.Wait()
}

func (h *Handler) sendToStream(
	i int,
	chunks []chunk,
	file *os.File,
	conn v1.Imaginator_DownloadFileByChunkServer,
) error {
	buf := make([]byte, chunks[i].bufSize)

	_, err := file.ReadAt(buf, int64(chunks[i].offset))
	if err != nil {
		return fmt.Errorf("file.ReadAt: %w", err)
	}

	err = conn.Send(&v1.DownloadFileResponse{
		Data:   buf,
		Offset: int64(i),
	})
	if err != nil {
		return fmt.Errorf("conn.Send: %w", err)
	}
	return nil
}
