package handler

import (
	"github.com/gin-gonic/gin"
	"linechat/services"
	"linechat/utils"
)

type FileServeHandler struct {
}

func NewFileServHandler() *FileServeHandler {
	return &FileServeHandler{}
}

func (h *FileServeHandler) GetFile(c *gin.Context) {
	utils.ServeFile(c)
	//c.JSON(200, gin.H{"file": "file"})
}
func (h *FileServeHandler) PostFile(c *gin.Context) {

	event := services.EventImpl{}

	//id := c.Param("id")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// open file
	files := form.File["banners"]
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer f.Close()
		// return image url
		url, err := utils.UploadFile(f, file.Filename)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		event.Banners = append(event.Banners, url)

	}

	c.JSON(200, gin.H{"file": event.Banners})
}
func (h *FileServeHandler) DeleteFile(c *gin.Context) {
	c.JSON(200, gin.H{"file": "file"})
}
