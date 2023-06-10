package disk

import (
	"fmt"
	"github.com/muhammadkhon-abdulloev/imaginator/internal/config"
	"github.com/muhammadkhon-abdulloev/imaginator/pkg/crypto"
	"github.com/muhammadkhon-abdulloev/imaginator/pkg/storage"
	"go.uber.org/fx"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	fileNameTemp = "%d_%s"
	maxLimit     = 100
)

type Storage struct {
	filesPath string
}

var (
	_      storage.StorageManager = (*Storage)(nil)
	Option                        = fx.Provide(New)
)

func New(cfg *config.Config) *Storage {
	return &Storage{
		filesPath: cfg.Server.ImagesPath,
	}
}
func (s *Storage) Upload(filename string, data []byte) (*storage.File, error) {
	filename = strings.ReplaceAll(filename, "_", "-")
	filename = strings.ReplaceAll(filename, "/", "-")

	file := storage.NewFile(fmt.Sprintf(fileNameTemp, time.Now().UnixMilli(), filename), 0, time.Now(), time.Now())

	newFile, err := os.Create(s.filesPath + "/" + file.Name)
	if err != nil {
		return nil, fmt.Errorf("os.Create: %w", err)
	}

	defer newFile.Close()

	size, err := newFile.Write(data)
	if err != nil {
		return nil, fmt.Errorf("newFile.Write: %w", err)
	}

	file.Size = int64(size)
	return file, nil
}

func (s *Storage) Download(filename string) (*storage.File, error) {
	file, err := os.Open(s.filesPath + "/" + filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll; %w", err)
	}

	checksum, err := crypto.FileChecksumSHA256(file)
	if err != nil {
		return nil, fmt.Errorf("crypto.FileChecksumSHA256: %w", err)
	}

	f := storage.NewEmptyFile()
	f.SetBytes(data)
	f.SetCheckSum(checksum)

	return f, nil
}

func (s *Storage) GetFile(filename string) (*os.File, error) {
	file, err := os.Open(s.filesPath + "/" + filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	return file, nil
}

func (s *Storage) ListAllFiles() ([]*storage.File, error) {
	var files []*storage.File

	dirEntries, err := os.ReadDir(s.filesPath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadDir: %w", err)
	}

	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			file, err := dirEntry.Info()
			if err != nil {
				return nil, fmt.Errorf("dirEntry.Info: %w", err)
			}

			createdAt, err := strconv.ParseInt(strings.Split(file.Name(), "_")[0], 10, 64)
			if err != nil {
				createdAt = file.ModTime().UnixMilli()
			}

			files = append(files, storage.NewFile(file.Name(), file.Size(), time.UnixMilli(createdAt), file.ModTime()))
		}
	}

	return files, nil
}

func (s *Storage) Paginate(limit, offset int64, files []*storage.File) []*storage.File {
	total := int64(len(files))
	if total < 1 {
		return files
	}

	if offset != 0 {
		offset--
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	firstEntry := offset * limit
	lastEntry := firstEntry + limit

	if firstEntry > total {
		firstEntry = total
	}

	if lastEntry > total {
		lastEntry = total
	}

	return files[firstEntry:lastEntry]
}
