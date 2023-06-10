package storage

import (
	"time"
)

type File struct {
	Name       string
	Size       int64
	CreatedAt  time.Time
	LastUpdate time.Time
	Bytes      []byte
	CheckSum   string
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

func NewEmptyFile() *File {
	return &File{}
}

func (f *File) SetName(name string) {
	f.Name = name
}

func (f *File) SetSize(size int64) {
	f.Size = size
}

func (f *File) SetCreatedAt(createdAt time.Time) {
	f.CreatedAt = createdAt
}

func (f *File) SetLastUpdate(lastUpdate time.Time) {
	f.LastUpdate = lastUpdate
}

func (f *File) SetBytes(b []byte) {
	f.Bytes = b
}

func (f *File) SetCheckSum(checksum string) {
	f.CheckSum = checksum
}
