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
	router.DELETE("/:id", r.appSettingHandler.DeleteAppSetting)
	router.GET("/:id", r.appSettingHandler.GetAppSetting)
	router.GET("/:id/member-types", r.appSettingHandler.GetMemberType)
	router.POST("/:id/member-types", r.appSettingHandler.PostMemberType)
	router.PUT("/:id/member-types", r.appSettingHandler.PutMemberType)
	router.DELETE("/:id/member-types/:memberTypeId", r.appSettingHandler.DeleteMemberType)
	router.GET("/:id/courses", r.appSettingHandler.GetCourses)
	router.POST("/:id/courses", r.appSettingHandler.PostAddCourse)
	router.PUT("/:id/courses", r.appSettingHandler.PutUpdateCourse)
	router.DELETE("/:id/courses/:courseId", r.appSettingHandler.DeleteCourse)

}
