package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	basePath string
	baseURL  string
}

func NewLocalStorage(basePath, baseURL string) *LocalStorage {
	return &LocalStorage{basePath: basePath, baseURL: baseURL}
}

func (ls *LocalStorage) Upload(ctx context.Context, input UploadInput, subDir string) (string, error) {
	dirPath := filepath.Join(ls.basePath, subDir)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	fullPath := filepath.Join(dirPath, input.Name)

	file, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, input.File); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return fmt.Sprintf("%s/%s/%s", ls.baseURL, subDir, input.Name), nil
}
