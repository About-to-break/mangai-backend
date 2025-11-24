package mangalocal

import (
	"backend/internal/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"path/filepath"
)

type UploadController struct {
	Storage  services.StorageService
	Queue    services.QueueService
	Bucket   string
	Exchange string
	Key      string
}

// IndexView отдаёт статический frontend
// @Summary Отдать главную страницу локального приложения
// @Description Возвращает index.html
// @Tags LocalManga
// @Produce html
// @Success 200 {string} string "index.html"
// @Router /local/ [get]
func IndexView(c *gin.Context) {
	c.File(filepath.Join("frontend", "index.html"))
}

// UploadView обрабатывает загрузку изображения
// @Summary Upload an image
// @Description Uploads an image file to the server
// @Tags LocalManga
// @Accept multipart/form-data
// @Param file formData file true "Image file"
// @Param make_dir formData string false "Whether to create a sub dir"
// @Produce json
// @Success 200 {object} map[string]string "File uploaded"
// @Failure 400 {object} map[string]string "No file received"
// @Router /local/upload [post]
func (u *UploadController) UploadView(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	src, err := file.Open()

	if err != nil {
		slog.Error("Error opening file", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	// UUID
	id := uuid.New().String()
	ext := filepath.Ext(file.Filename)
	makeDir := c.DefaultPostForm("make_dir", "true")
	var objectName string
	if makeDir == "true" {
		objectName = fmt.Sprintf("%s/%s%s", id, id, ext)
	} else {
		objectName = fmt.Sprintf("%s%s", id, ext)
	}

	// Storage
	err = u.Storage.UploadFile(u.Bucket, objectName, src, file.Size, file.Header.Get("Content-Type"))

	if err != nil {
		slog.Error("Upload Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	slog.Debug("Successfully uploaded file")
	// Queue
	if u.Queue != nil && u.Exchange != "" && u.Key != "" {
		msg := map[string]string{"uuid": id, "object_name": objectName}
		body, _ := json.Marshal(msg)
		err = u.Queue.Publish(u.Exchange, u.Key, body)
		if err != nil {
			slog.Error("Error sending message to queue")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		slog.Debug("Successfully set queue task")
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "uploaded",
		"filename": file.Filename,
	})

}
