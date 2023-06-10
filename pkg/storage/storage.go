package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/muhammadkhon-abdulloev/imaginator/pkg/crypto"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	fileNameTemp = "%d_%s"
	//bufSize      = 500 * 1024
)

var _ IStorage = (*Storage)(nil)

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

	defer newFile.Close()

	size, err := newFile.Write(data)
	if err != nil {
		return nil, fmt.Errorf("newFile.Write: %w", err)
	}

	file.Size = int64(size)
	return file, nil
}

func (s *Storage) Download(filename string) (*File, error) {
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

	f := NewEmptyFile()
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
