package handler

import (
	"github.com/gin-gonic/gin"
	"linechat/services"
)

type MemberHandler struct {
	memberService services.MemberService
}

func NewMemberHandler(memberService services.MemberService) *MemberHandler {
	return &MemberHandler{
		memberService: memberService,
	}
}

func (h *MemberHandler) GetMembers(c *gin.Context) {
	result, err := h.memberService.GetMembers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if result == nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.JSON(200, gin.H{"members": result})
}
func (h *MemberHandler) GetMember(c *gin.Context) {
	linedId := c.Param("linedId")
	if linedId == "" {
		c.JSON(400, gin.H{"error": "Invalid linedId"})
		return
	}
	result, er := h.memberService.GetMemberByLineId(linedId)
	if er != nil {
		c.JSON(500, gin.H{"error": er.Error()})
		return
	}
	if result == nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.JSON(200, gin.H{"member": result})

}
