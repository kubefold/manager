package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kubefold/manager/internal/dto"
	"github.com/sirupsen/logrus"
)

type UploadService interface {
	UploadArtifacts(bucket string) error
}

type uploadService struct {
	config dto.Config
}

func newUploadService(config dto.Config) UploadService {
	return &uploadService{
		config: config,
	}
}

func (u uploadService) UploadArtifacts(bucket string) error {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	return filepath.Walk(u.config.OutputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(u.config.OutputPath, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		s3Key := strings.ReplaceAll(relPath, "\\", "/")

		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", path, err)
		}
		defer file.Close()

		logrus.Infof("Uploading %s to s3://%s/%s", path, bucket, s3Key)
		_, err = client.PutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(s3Key),
			Body:   file,
		})
		if err != nil {
			return fmt.Errorf("failed to upload file %s to S3: %w", path, err)
		}

		return nil
	})
}
