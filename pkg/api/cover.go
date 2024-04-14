package api

import (
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go/v7"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync-in-between-s3/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CoverResponse If you put lower-case property, it won't be returned, because it's `private`
type CoverResponse struct {
	Response string `json:"response"`
}

// @BasePath /api

// Cover godoc
// @Summary check if cover available, if not - download from local and upload to remote
// @Schemes
// @Param type query string true "Type"
// @Param id query int true "ID"
// @Param l_bucket query string true "Local bucket"
// @Param r_bucket query string true "Remote bucket"
// @Description Check if cover available, if not - download from local and upload to remote
// @Tags api
// @Accept json
// @Produce json
// @Success 200 {object} api.CoverResponse
// @Router /cover [get]
func Cover(c *gin.Context) {
	entityType := c.Query("type")
	entityID := c.Query("id")

	localBucket := c.Query("l_bucket")
	remoteBucket := c.Query("r_bucket")

	res := c.Writer

	if entityType == "" || !(entityType == "anime" || entityType == "manga") {
		http.Error(res, "Empty entityType provided", http.StatusBadRequest)
		return
	}
	entityIDInt, err := strconv.Atoi(entityID)
	if entityID == "" || err != nil {
		http.Error(res, "Empty entityID provided", http.StatusBadRequest)
		return
	}
	response := CoverResponse{Response: entityType}

	ok, err := http.Get(utils.GetUrl(entityType, entityIDInt))
	fmt.Println("Status: ", ok.StatusCode)

	if err != nil || ok.StatusCode >= 400 {
		fmt.Println("Process: " + entityID)
		minioClient, err := utils.GetLocalMinio(c)
		if err != nil {
			http.Error(res, "Cannot connect to MinIO: "+err.Error(), http.StatusInternalServerError)
			return
		}

		yandexClient, err := utils.GetRemoteMinio(c)
		if err != nil {
			http.Error(res, "Cannot connect to Yandex: "+err.Error(), http.StatusInternalServerError)
			return
		}

		dir, err := os.MkdirTemp("", entityType)
		if err != nil {
			http.Error(res, "Cannot create dir: "+err.Error(), http.StatusInternalServerError)
			return
		}

		defer func(path string) {
			err := os.RemoveAll(path)
			if err != nil {
				http.Error(res, "Failed to clean /tmp: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}(dir)

		minioPrefix := fmt.Sprintf("covers/shikimori/%s/", entityType)
		yandexPrefix := fmt.Sprintf("%s/covers/", entityType)

		objectKey := fmt.Sprintf("%d.jpg", entityIDInt)
		path := filepath.Join(dir, objectKey)

		err = minioClient.FGetObject(c, localBucket, filepath.Join(minioPrefix, objectKey), path, minio.GetObjectOptions{})
		if err != nil {
			http.Error(res, "Cant download: "+err.Error(), http.StatusInternalServerError)
			return
		}

		info, err := yandexClient.FPutObject(c, remoteBucket, filepath.Join(yandexPrefix, objectKey), path, minio.PutObjectOptions{})
		if err != nil {
			http.Error(res, "Cant upload: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(info.Key)
	} else {
		fmt.Println(fmt.Sprintf("Skip item id=%s & type=%s", entityID, entityType))
	}

	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		http.Error(res, "Bad response value: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
