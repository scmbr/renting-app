package cloud

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)
	


func NewS3Client() (*s3.Client, error) {
    // Подгружаем конфигурацию из ~/.aws/*
    cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("ru-central1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("YC_ACCESS_KEY"),
			os.Getenv("YC_SECRET_KEY"),
            "",
		)),
	)
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
        // Указываем базовый эндпоинт Yandex
        o.BaseEndpoint = aws.String("https://storage.yandexcloud.net")
        o.RequestChecksumCalculation = aws.RequestChecksumCalculationWhenRequired
        // Опционально: кастомный резолвер (для сложных сценариев)
        o.EndpointResolverV2 = s3.NewDefaultEndpointResolverV2()
    }), nil
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
func UploadAvatarToS3(fileHeader *multipart.FileHeader) (string, error) {
    prefix := "avatars/"
    bucketName:="profile-pictures"
    s3Client, err := NewS3Client()
    if err != nil {
        return "", err
    } 
    err=TestS3Connection(s3Client,bucketName)
    if err != nil {
        return "", err
    } 
   

    url, err := uploadToS3(s3Client, fileHeader, generateFilename(fileHeader.Filename),bucketName,prefix)
    return url, err
}
func generateFilename(original string) string {
    ext := filepath.Ext(original)
    return fmt.Sprintf("%s%s", uuid.New().String(), ext)
}

func uploadToS3(s3Client *s3.Client, fileHeader *multipart.FileHeader, filename string, bucketName string, prefix string) (string, error) {
    file, err := fileHeader.Open()
    if err != nil {
        return "", err
    }
    defer file.Close()

    
    // Загружаем файл в S3
    _, err = s3Client.PutObject(context.Background(), &s3.PutObjectInput{
        Bucket:      aws.String(bucketName),
        Key:         aws.String(prefix + filename),
        Body:        file,
        ContentType: aws.String("image/png"), // Укажите MIME-тип
        ContentLength: nil,
    })
    if err != nil {
        return "", err
    }

    return "https://storage.yandexcloud.net/" + bucketName + "/" + prefix + filename, nil
}


