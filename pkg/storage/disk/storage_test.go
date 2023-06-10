package diskstorage

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestStorage_Upload(t *testing.T) {
	storage := NewStorage(".")

	testCases := []struct {
		filename string
		data     []byte
	}{
		{
			filename: "Test1.txt",
			data:     []byte("Hello"),
		},
		{
			filename: "Test1.txt",
			data:     []byte("Hello"),
		},
	}

	for _, testCase := range testCases {
		file, err := storage.Upload(testCase.filename, testCase.data)
		require.Nil(t, err)

		data, err := storage.Download(file.Name)
		require.Nil(t, err)

		require.Equal(t, data.Bytes, testCase.data)
	}
}

func TestStorage_Download(t *testing.T) {
	storage := NewStorage(".")
	data := []byte("hello")

	filename, err := storage.Upload("file1.txt", data)
	require.Nil(t, err, "storage.Upload: want nil error, got err: %s", err)

	testCases := []struct {
		filename string
		data     []byte
		err      error
	}{
		{
			filename: "Test1.txt",
			err:      os.ErrNotExist,
		},
		{
			filename: filename.Name,
			data:     data,
		},
	}

	for _, testCase := range testCases {
		data, err := storage.Download(testCase.filename)
		if testCase.err != nil {
			require.ErrorIs(t, err, testCase.err)
			continue
		}

		require.Equal(t, data.Bytes, testCase.data)
	}
}

func TestStorage_ListAllFiles(t *testing.T) {
	storage := NewStorage(".")
	files, err := storage.ListAllFiles()
	require.Nil(t, err)
	require.NotNil(t, files)
}
