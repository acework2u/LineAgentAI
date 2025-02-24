package handler

import (
	"github.com/gin-gonic/gin"
	"linechat/services"
	"linechat/utils"
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
	c.JSON(200, gin.H{"memberType": "ok"})
}
func (h *AppSettingHandler) PostMemberType(c *gin.Context) {
	memberType := services.MemberTypeImpl{}
	err := c.ShouldBindJSON(&memberType)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	c.JSON(200, gin.H{"memberType": "ok"})

}
