package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"sync-in-between-s3/pkg/config"
)

func GetLocalClient(cfg *config.Environment) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.LocalEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.LocalAccessKeyID, cfg.LocalSecretKey, ""),
		Secure: cfg.LocalUseSSL,
	})

	return minioClient, err
}

func GetRemoteClient(cfg *config.Environment) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.RemoteEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.RemoteAccessKeyID, cfg.RemoteSecretKey, ""),
		Secure: cfg.RemoteUseSSL,
	})

	return minioClient, err
}
