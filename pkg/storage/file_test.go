package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {
	tt := struct {
		filename   string
		filesize   int64
		createdAt  time.Time
		lastUpdate time.Time
	}{
		filename:   "SomeName",
		filesize:   10,
		createdAt:  time.Now(),
		lastUpdate: time.Now(),
	}

	file := NewFile(
		tt.filename,
		tt.filesize,
		tt.createdAt,
		tt.lastUpdate,
	)

	require.Equal(t, file.Name, tt.filename)
	require.Equal(t, file.Size, tt.filesize)
	require.Equal(t, file.CreatedAt, tt.createdAt)
	require.Equal(t, file.LastUpdate, tt.lastUpdate)

}
