package storage

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const fileNameTemp = "%d_%s"

type Storage struct {
	filesPath string
}

func NewStorage(filesPath string) *Storage {
	return &Storage{
		filesPath: filesPath,
	}
}
func (s *Storage) Upload(filename string, data []byte) (*File, error) {
	filename = strings.ReplaceAll(filename, "_", "-")
	filename = strings.ReplaceAll(filename, "/", "-")

	file := NewFile(fmt.Sprintf(fileNameTemp, time.Now().UnixMilli(), filename), 0, time.Now(), time.Now())
	if file.Name == "" {
		file.Name = fmt.Sprintf(fileNameTemp, file.CreatedAt.UnixMilli(), uuid.NewString())
	}

	newFile, err := os.Create(s.filesPath + "/" + file.Name)
	if err != nil {
		return nil, fmt.Errorf("os.Create: %w", err)
	}

	size, err := newFile.Write(data)
	if err != nil {
		return nil, fmt.Errorf("newFile.Write: %w", err)
	}

	file.Size = int64(size)
	return file, nil
}

func (s *Storage) Download(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll; %w", err)
	}

	return data, nil
}

func (s *Storage) ListAllFiles() ([]*File, error) {
	var files []*File

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

			files = append(files, NewFile(file.Name(), file.Size(), time.UnixMilli(createdAt), file.ModTime()))
		}
	}

	return files, nil
}
