package storage

import "time"

type File struct {
	Name       string
	Size       int64
	CreatedAt  time.Time
	LastUpdate time.Time
}

func NewFile(
	name string,
	size int64,
	createdAt time.Time,
	lastUpdate time.Time,
) *File {
	return &File{
		Name:       name,
		Size:       size,
		CreatedAt:  createdAt,
		LastUpdate: lastUpdate,
	}
}
