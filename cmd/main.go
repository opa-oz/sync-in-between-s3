package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "sync-in-between-s3/cmd/docs"
	"sync-in-between-s3/pkg/api"
	"sync-in-between-s3/pkg/config"
	"sync-in-between-s3/pkg/middlewares"
	"sync-in-between-s3/pkg/minio"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.Production {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	r.Use(middlewares.RequestLogger())
	r.Use(middlewares.ResponseLogger())

	localClient, err := minio.GetLocalClient(cfg)
	remoteClient, err := minio.GetRemoteClient(cfg)

	r.GET("/healz", api.Healz)
	r.GET("/ready", api.Ready)
	// {@link https://github.com/swaggo/gin-swagger?tab=readme-ov-file}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	rg := r.Group("/api")
	rg.Use(middlewares.CfgMiddleware(cfg))
	rg.Use(middlewares.LocalMinioMiddleware(localClient))
	rg.Use(middlewares.RemoteMinioMiddleware(remoteClient))

	{
		rg.GET("/cover", api.Cover)
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	err = r.Run(port)
	if err != nil {
		fmt.Println(err)
		return
	}
}
