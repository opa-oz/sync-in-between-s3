package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func LocalMinioMiddleware(minioClient *minio.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("LocalMinio", minioClient)
		c.Next()
	}
}

func RemoteMinioMiddleware(minioClient *minio.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("RemoteMinio", minioClient)
		c.Next()
	}
}
