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
	website  string
}

func NewFileStorage(client *s3.Client, bucket, endpoint string, website string) *FileStorage {
	return &FileStorage{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
		website:  website,
	}
}
func TestS3Connection(s3Client *s3.Client, bucketName string) error {
	ctx := context.Background()
	_, err := s3Client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return fmt.Errorf("failed to head bucket: %w", err)
	}

	fmt.Println("Connection successful. Bucket exists.")
	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("failed to head bucket: %w", err)
	}
	for _, bucket := range result.Buckets {
		fmt.Printf("bucket=%s creation time=%s\n", aws.ToString(bucket.Name), bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
	}
	return nil
}
func (fs *FileStorage) Upload(ctx context.Context, input UploadInput) (string, error) {
	filename := generateFilename(input.Name)
	_, err := fs.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:        aws.String(fs.bucket),
		Key:           aws.String(filename),
		Body:          input.File,
		ContentType:   aws.String(input.ContentType),
		ContentLength: nil,
	})
	if err != nil {
		return "", err
	}

	return fs.generateFileURL(filename), nil
}

func (fs *FileStorage) generateFileURL(filename string) string {
	return fmt.Sprintf("https://%s.%s/%s", fs.bucket, fs.website, filename)
}
func generateFilename(original string) string {
	ext := filepath.Ext(original)
	return fmt.Sprintf("%s%s", uuid.New().String(), ext)
}
