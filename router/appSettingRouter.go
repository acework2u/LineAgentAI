package router

import (
	"github.com/gin-gonic/gin"
	"linechat/handler"
)

type AppSettingRouter struct {
	appSettingHandler *handler.AppSettingHandler
}

func NewAppSettingRouter(appSettingHandler *handler.AppSettingHandler) *AppSettingRouter {
	return &AppSettingRouter{appSettingHandler: appSettingHandler}
}

func (r *AppSettingRouter) AppSettingRouter(rg *gin.RouterGroup) {
	router := rg.Group("/settings")

	router.GET("", r.appSettingHandler.GetAppSetting)
	router.POST("/", r.appSettingHandler.PostAppSetting)
	router.GET("/:id", r.appSettingHandler.GetAppSetting)
	router.GET("/:id/member-type", r.appSettingHandler.GetMemberType)
	router.POST("/:id/member-type", r.appSettingHandler.PostMemberType)

}
