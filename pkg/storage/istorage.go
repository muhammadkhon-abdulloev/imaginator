package storage

import "os"

type IStorage interface {
	Upload(filename string, data []byte) (*File, error)
	Download(filename string) (*File, error)
	GetFile(filename string) (*os.File, error)
	ListAllFiles() ([]*File, error)
}
