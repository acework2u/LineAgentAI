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
	router.GET("/:id/course-types", r.appSettingHandler.GetCourseTypeList)
	router.POST("/:id/course-types", r.appSettingHandler.PostAddCourseType)
	router.PUT("/:id/course-types", r.appSettingHandler.PutUpdateCourseType)
	router.DELETE("/:id/course-types/:courseTypeId", r.appSettingHandler.DeleteCourseType)
	router.GET("/:id/clinics", r.appSettingHandler.GetClinics)
	router.POST("/:id/clinics", r.appSettingHandler.PostAddClinic)
	router.PUT("/:id/clinics", r.appSettingHandler.PutUpdateClinic)
	router.DELETE("/:id/clinics/:clinicId", r.appSettingHandler.DeleteClinic)
}
