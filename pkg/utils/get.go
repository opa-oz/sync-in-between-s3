package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"sync-in-between-s3/pkg/config"
)

func GetLocalMinio(c *gin.Context) (*minio.Client, error) {
	r := c.Value("LocalMinio")

	if r == nil {
		err := fmt.Errorf("could not retrieve LocalMinio")
		return nil, err
	}

	rdb, ok := r.(*minio.Client)
	if !ok {
		err := fmt.Errorf("variable LocalMinio has wrong type")
		return nil, err
	}

	return rdb, nil
}

func GetRemoteMinio(c *gin.Context) (*minio.Client, error) {
	r := c.Value("RemoteMinio")

	if r == nil {
		err := fmt.Errorf("could not retrieve RemoteMinio")
		return nil, err
	}

	rdb, ok := r.(*minio.Client)
	if !ok {
		err := fmt.Errorf("variable RemoteMinio has wrong type")
		return nil, err
	}

	return rdb, nil
}

func GetConfig(c *gin.Context) (*config.Environment, error) {
	r := c.Value("Config")

	if r == nil {
		err := fmt.Errorf("could not retrieve Config")
		return nil, err
	}

	rdb, ok := r.(*config.Environment)
	if !ok {
		err := fmt.Errorf("variable Config has wrong type")
		return nil, err
	}

	return rdb, nil
}
