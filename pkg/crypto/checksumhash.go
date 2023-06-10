package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func FileChecksumSHA256(file *os.File) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", fmt.Errorf("io.Copy: %w", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func FileChecksumMD5(file *os.File) (string, error) {
	h := md5.New()

	if _, err := io.Copy(h, file); err != nil {
		return "", fmt.Errorf("io.Copy: %w", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
