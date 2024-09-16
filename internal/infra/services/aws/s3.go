package aws

import (
	"context"
	"fmt"
	"mime"
	"mime/multipart"
	"path/filepath"

	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/felipeversiane/donate-api/internal/infra/config"
)

type ObjectStorageInterface interface {
	CreateBucket(ctx context.Context) error
	UploadFile(ctx context.Context, bucketKey string, file multipart.File) (string, error)
	DeleteFile(ctx context.Context, key string) error
}

type objectStorage struct {
	client *s3.S3
	bucket string
	region string
	acl    string
	url    string
}

func NewObjectStorage(session *session.Session, config config.CloudServiceConfig) ObjectStorageInterface {
	return &objectStorage{
		client: s3.New(session),
		bucket: config.ObjectStorage.Bucket,
		region: config.Region,
		acl:    config.ObjectStorage.ACL,
		url:    config.ObjectStorage.URL,
	}
}

func (o *objectStorage) CreateBucket(ctx context.Context) error {
	_, err := o.client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(o.bucket),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(o.region),
		},
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == s3.ErrCodeBucketAlreadyOwnedByYou {
				slog.Warn("Bucket already owned by you", "bucket", o.bucket)
				return nil
			}
			if aerr.Code() == s3.ErrCodeBucketAlreadyExists {
				slog.Warn("Bucket already exists", "bucket", o.bucket)
				return nil
			}
		}
		slog.Error("Failed to create bucket", "bucket", o.bucket, "error", err)
		return fmt.Errorf("unable to create bucket %s: %v", o.bucket, err)
	}

	return nil
}

func (o *objectStorage) UploadFile(ctx context.Context, bucketKey string, file multipart.File) (string, error) {
	defer file.Close()

	ext := filepath.Ext(bucketKey)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := o.client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(o.bucket),
		Key:         aws.String(bucketKey),
		Body:        file,
		ACL:         aws.String(o.acl),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		slog.Error("Failed to upload file", "bucket", o.bucket, "bucketKey", bucketKey, "error", err)
		return "", fmt.Errorf("unable to upload file %s to bucket %s: %v", bucketKey, o.bucket, err)
	}

	url := fmt.Sprintf("%s/%s/%s", o.url, o.bucket, bucketKey)
	return url, nil
}

func (o *objectStorage) DeleteFile(ctx context.Context, key string) error {
	_, err := o.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(o.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		slog.Error("Failed to delete file", "bucket", o.bucket, "key", key, "error", err)
		return fmt.Errorf("error deleting file %s from bucket %s: %v", key, o.bucket, err)
	}

	return nil
}
