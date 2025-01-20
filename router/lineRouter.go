package router

import (
	"github.com/gin-gonic/gin"
	"linechat/handler"
)

type LineRouter struct {
	lineHandler *handler.LineWebhookHandler
}

func NewLineRouter(lineHandler *handler.LineWebhookHandler) *LineRouter {
	return &LineRouter{
		lineHandler: lineHandler,
	}
}
func (r *LineRouter) LineHookRouter(rg *gin.RouterGroup) {
	rg.POST("/linehook", r.lineHandler.LineHookHandle)

}
