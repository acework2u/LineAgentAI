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
	rg.POST("/callback", r.lineHandler.LineCallback)
	rg.POST("/register", r.lineHandler.LineRegister)
	rg.POST("/login", r.lineHandler.LineLogin)
	rg.POST("/logout", r.lineHandler.LineLogout)
	rg.POST("/chat", r.lineHandler.LineChat)
	rg.GET("/profile", r.lineHandler.GetLineProfile)
	rg.PUT("/profile", r.lineHandler.PutLineProfile)

}
