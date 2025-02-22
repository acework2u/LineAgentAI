package handler

import (
	"github.com/gin-gonic/gin"
	"linechat/services"
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
