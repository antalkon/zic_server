package storage

import (
	"backend/pkg/config"
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOClient struct {
	Client *minio.Client
	Bucket string
}

func NewMinIOClient(cfg *config.Config) (*MinIOClient, error) {
	client, err := minio.New(cfg.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		Secure: cfg.S3UseSSL,
	})
	if err != nil {
		log.Fatalf("Failed to connect to MinIO: %v", err)
		return nil, err
	}

	log.Println("Connected to MinIO successfully")

	exists, err := client.BucketExists(context.Background(), cfg.S3Bucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		log.Printf("Bucket %s does not exist, creating...", cfg.S3Bucket)
		err = client.MakeBucket(context.Background(), cfg.S3Bucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
			return nil, err
		}
	}

	return &MinIOClient{Client: client, Bucket: cfg.S3Bucket}, nil
}

func (m *MinIOClient) UploadFile(objectName, filePath, contentType string) error {
	ctx := context.Background()
	_, err := m.Client.FPutObject(ctx, m.Bucket, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Printf("Failed to upload file %s: %v", objectName, err)
		return err
	}
	log.Printf("File %s uploaded successfully", objectName)
	return nil
}

// GetFileURL получает публичный URL файла
func (m *MinIOClient) GetFileURL(objectName string) (string, error) {
	url, err := m.Client.PresignedGetObject(context.Background(), m.Bucket, objectName, 0, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
