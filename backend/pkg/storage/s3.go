package storage

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type FileStorage struct {
	client   *s3.Client
	bucket   string
	endpoint string
}

func NewFileStorage(client *s3.Client, bucket, endpoint string) *FileStorage {
	return &FileStorage{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
	}
}

func (fs *FileStorage) Upload(ctx context.Context, input UploadInput) (string, error) {
	filename := generateFilename(input.Name)
	_, err := fs.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:        aws.String(fs.bucket),
		Key:           aws.String(filename),
		Body:          input.File,
		ContentType:   aws.String("image/png"), // Укажите MIME-тип
		ContentLength: nil,
	})
	if err != nil {
		return "", err
	}

	return fs.generateFileURL(input.Name), nil
}

func (fs *FileStorage) generateFileURL(filename string) string {
	return fmt.Sprintf("https://%s.%s/%s", fs.bucket, fs.endpoint, filename)
}
func generateFilename(original string) string {
	ext := filepath.Ext(original)
	return fmt.Sprintf("%s%s", uuid.New().String(), ext)
}
