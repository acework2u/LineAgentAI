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

	rg.GET("/profile", r.lineHandler.GetLineProfile)
	rg.GET("/check-join-event", r.lineHandler.GetCheckEventJoin)
	rg.GET("/join-event", r.lineHandler.GetEventJoin)
	rg.POST("/event-check-in", r.lineHandler.PostCheckInEvent)
	rg.POST("/my-event", r.lineHandler.PostMyEvents)
	rg.POST("/linehook", r.lineHandler.LineHookHandle)
	rg.POST("/webhook", r.lineHandler.LineWebhook)
	rg.POST("/callback", r.lineHandler.LineCallback)
	rg.POST("/register", r.lineHandler.LineRegister)
	rg.POST("/login", r.lineHandler.LineLogin)
	rg.POST("/logout", r.lineHandler.LineLogout)
	rg.POST("/chat", r.lineHandler.LineChat)
	rg.POST("/check-registration", r.lineHandler.CheckMemberRegister)
	rg.POST("/member-update", r.lineHandler.PostUpdateMember)
	rg.POST("/join-event", r.lineHandler.PostJoinEvent)
	rg.PUT("/profile", r.lineHandler.PutLineProfile)

}
