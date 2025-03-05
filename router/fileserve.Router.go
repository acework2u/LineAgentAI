package router

import (
	"github.com/gin-gonic/gin"
	"linechat/handler"
)

type FileServeRouter struct {
	fileServeHandler *handler.FileServeHandler
}

func NewFileServeRouter(serveHandler *handler.FileServeHandler) *FileServeRouter {
	return &FileServeRouter{
		fileServeHandler: serveHandler,
	}
}

func (r *FileServeRouter) FileServeRouter(rg *gin.RouterGroup) {
	router := rg.Group("/files")

	router.GET("/:id", r.fileServeHandler.GetFile)
	router.POST("/", r.fileServeHandler.PostFile)
	router.POST("", r.fileServeHandler.PostFile)
}
