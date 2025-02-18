package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linechat/services"
	"linechat/utils"
)

type StaffHandler struct {
	staffService services.StaffService
}

func NewStaffHandler(staffService services.StaffService) *StaffHandler {
	return &StaffHandler{
		staffService: staffService,
	}
}

func (s *StaffHandler) GetStaffs(c *gin.Context) {
	rsult, err := s.staffService.GetStaffs()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": rsult})
}
func (s *StaffHandler) GetStaff(c *gin.Context) {
	userId := c.Param("userId")
	rsult, err := s.staffService.GetStaffById(userId)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": rsult})
}
func (s *StaffHandler) CreateStaff(c *gin.Context) {
	staffInfo := services.StaffRegister{}
	err := c.ShouldBindJSON(&staffInfo)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	// Staff Register
	err = s.staffService.Register(&staffInfo)
	if err != nil {

		//c.JSON(400, gin.H{"message": err.Error()})
		c.JSON(400, gin.H{"message": "Duplicate key error: the email and name combination already exists"})
		return
	}

	c.JSON(200, gin.H{"message": "register staff successfully"})

}
func (s *StaffHandler) LoginStaff(c *gin.Context) {
	staffLogin := services.StaffLogin{}
	err := c.ShouldBindJSON(&staffLogin)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}

	fmt.Println(staffLogin)

	// login service
	token, err := s.staffService.LoginStaff(&staffLogin)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	// response data token
	c.JSON(200, gin.H{"token": token})

}
