package storage

import (
	"os"
)

type StorageManager interface {
	Upload(filename string, data []byte) (*File, error)
	Download(filename string) (*File, error)
	GetFile(filename string) (*os.File, error)
	ListAllFiles() ([]*File, error)
	Paginate(limit, offset int64, files []*File) []*File
}
