package handler

import (
	"fmt"
	v1 "github.com/muhammadkhon-abdulloev/imaginator/gen/go/imaginator/v1"
	"golang.org/x/sync/errgroup"
	"os"
)

type chunk struct {
	bufSize int
	offset  int
}

func (h *Handler) streamFile(
	filename string,
	chunkBuffSize int,
	conn v1.Imaginator_DownloadFileByChunkServer,
) error {
	file, err := h._storage.DownloadByChunks(filename)
	if err != nil {
		err = fmt.Errorf("h._storage.Download: %w", err)
		h.lg.Error(err)
		return err
	}

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

		if i == chunksQ-1 {
			chunks[i].bufSize = filesize % chunkBuffSize
		}
	}

	eg := &errgroup.Group{}
	for i := 0; i < chunksQ; i++ {
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

	}

	err = conn.Send(&v1.DownloadFileResponse{
		Data: buf,
	})
	if err != nil {
		err = fmt.Errorf("conn.Send: %w", err)
	}
	return nil
}
