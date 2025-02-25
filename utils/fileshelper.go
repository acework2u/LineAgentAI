package utils

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io"
	"linechat/conf"
	"net/http"
)

func UploadFile(file io.Reader, filename string) (string, error) {
	//Upload file to GridFS and return URL
	appConfig, _ := conf.NewAppConfig()
	client := conf.ConnectionDB()
	bucket, err := gridfs.NewBucket(client.Database(appConfig.Db.DbName), nil)
	if err != nil {
		return "", err
	}
	uploadStream, err := bucket.OpenUploadStream(filename)
	if err != nil {
		return "", err
	}
	defer uploadStream.Close()
	_, err = io.Copy(uploadStream, file)
	if err != nil {
		return "", err
	}
	fileID := uploadStream.FileID.(primitive.ObjectID).Hex()
	return "api/v1/files/" + fileID, nil
}
func ServeFile(c *gin.Context) {
	fileID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	appConfig, _ := conf.NewAppConfig()
	client := conf.ConnectionDB()
	bucket, err := gridfs.NewBucket(client.Database(appConfig.Db.DbName), nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	downloadStream, err := bucket.OpenDownloadStream(objID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer downloadStream.Close()

	c.Header("Content-Type", "image/jpeg") // Adjust MIME type dynamically if needed
	_, err = io.Copy(c.Writer, downloadStream)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
