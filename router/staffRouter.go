package router

import (
	"github.com/gin-gonic/gin"
	"linechat/handler"
	"linechat/middleware"
)

type StaffRouter struct {
	staffHandler *handler.StaffHandler
}

func NewStaffRouter(staffHandler *handler.StaffHandler) *StaffRouter {
	return &StaffRouter{staffHandler: staffHandler}
}

func (r *StaffRouter) StaffRouter(rg *gin.RouterGroup) {

	rt1 := rg.Group("/auth")
	rt1.GET("/login", r.staffHandler.LoginStaff)

	rt := rg.Group("/staffs")

	rt.GET("", middleware.AuthMiddleware(), r.staffHandler.GetStaffs)
	rt.GET("/:staffId", middleware.AuthMiddleware(), r.staffHandler.GetStaff)
	rt.POST("/", r.staffHandler.CreateStaff)
	rt.POST("/login", r.staffHandler.LoginStaff)

}
