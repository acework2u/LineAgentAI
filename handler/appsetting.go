package handler

import (
	"github.com/gin-gonic/gin"
	"linechat/services"
	"linechat/utils"
	"log"
)

type AppSettingHandler struct {
	appSettingServ services.AppSettingsService
}

func NewAppSettingHandler(appSettingServ services.AppSettingsService) *AppSettingHandler {
	return &AppSettingHandler{
		appSettingServ: appSettingServ,
	}
}

func (h *AppSettingHandler) GetAppSetting(c *gin.Context) {
	appSettingInfo, err := h.appSettingServ.GetAppSettings()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"appSetting": appSettingInfo})
}

func (h *AppSettingHandler) PostAppSetting(c *gin.Context) {

	setttingInfo := services.AppSettings{}
	err := c.ShouldBindJSON(&setttingInfo)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.CreateAppSettings(&setttingInfo)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"appSetting": "ok"})
}

func (h *AppSettingHandler) GetMemberType(c *gin.Context) {
	id := c.Param("id")
	memberType, err := h.appSettingServ.MemberTypesList(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"memberType": memberType})
}
func (h *AppSettingHandler) PostMemberType(c *gin.Context) {
	appIds := c.Param("id")
	memberType := services.MemberTypeImpl{}
	log.Println(appIds)
	log.Println(memberType)
	log.Println("post memberType")

	err := c.ShouldBindJSON(&memberType)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.AddMemberType(appIds, &memberType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "add memberType successfully"})

}
func (h *AppSettingHandler) PutMemberType(c *gin.Context) {
	appIds := c.Param("id")
	memberType := services.MemberTypeImpl{}
	err := c.ShouldBindJSON(&memberType)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
	}
	err = h.appSettingServ.UpdateMemberType(appIds, &memberType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "update memberType successfully"})
}
func (h *AppSettingHandler) DeleteMemberType(c *gin.Context) {
	appIds := c.Param("id")
	memberType := services.MemberTypeImpl{}
	err := c.ShouldBindJSON(&memberType)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
	}
	err = h.appSettingServ.DeleteMemberType(appIds, &memberType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "delete memberType successfully"})
}

func (h *AppSettingHandler) DeleteAppSetting(c *gin.Context) {
	appIds := c.Param("id")
	if appIds == "" {
		c.JSON(400, gin.H{"error": "Invalid appIds"})
		return
	}
	err := h.appSettingServ.DeleteAppSettings(appIds)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "delete appSetting successfully"})
}
