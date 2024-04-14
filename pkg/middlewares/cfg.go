package middlewares

import (
	"github.com/gin-gonic/gin"
	"sync-in-between-s3/pkg/config"
)

func CfgMiddleware(cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Config", cfg)
		c.Next()
	}
}
