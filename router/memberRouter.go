package router

import (
	"github.com/gin-gonic/gin"
	"linechat/handler"
)

type MemberRouter struct {
	memberHandler *handler.MemberHandler
}

func NewMemberRouter(memberHandler *handler.MemberHandler) *MemberRouter {
	return &MemberRouter{memberHandler: memberHandler}
}

func (r *MemberRouter) MemberRouter(rg *gin.RouterGroup) {
	router := rg.Group("/members")
	router.GET("", r.memberHandler.GetMembers)
	router.GET("/:id", r.memberHandler.GetMember)
}
