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
func (h *AppSettingHandler) GetCourses(c *gin.Context) {
	appIds := c.Param("id")
	if appIds == "" {
		c.JSON(400, gin.H{"error": "Invalid app id"})
		return
	}
	courses, err := h.appSettingServ.CourseList(appIds)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": courses})
}
func (h *AppSettingHandler) PostAddCourse(c *gin.Context) {
	appId := c.Param("id")
	course := services.Course{}
	err := c.ShouldBindJSON(&course)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.AddCourse(appId, &course)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "add course successfully"})
}
func (h *AppSettingHandler) PutUpdateCourse(c *gin.Context) {
	appId := c.Param("id")
	course := services.Course{}
	err := c.ShouldBindJSON(&course)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.UpdateCourse(appId, &course)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Update course successfully"})
}
func (h *AppSettingHandler) DeleteCourse(c *gin.Context) {
	appId := c.Param("id")
	course := services.Course{}
	err := c.ShouldBindJSON(&course)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.DeleteCourse(appId, &course)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "delete course successfully"})
}
func (h *AppSettingHandler) GetCourseTypeList(c *gin.Context) {
	appIds := c.Param("id")
	if appIds == "" {
		c.JSON(400, gin.H{"error": "Invalid app id"})
		return
	}
	courseTypes, err := h.appSettingServ.CourseTypeList(appIds)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": courseTypes})
}
func (h *AppSettingHandler) PostAddCourseType(c *gin.Context) {
	appId := c.Param("id")
	courseType := services.CourseType{}
	err := c.ShouldBindJSON(&courseType)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.AddCourseType(appId, &courseType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "add course type successfully"})
}
func (h *AppSettingHandler) PutUpdateCourseType(c *gin.Context) {
	appId := c.Param("id")
	courseType := services.CourseType{}
	err := c.ShouldBindJSON(&courseType)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.UpdateCourseType(appId, &courseType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "update course type successfully"})
}
func (h *AppSettingHandler) DeleteCourseType(c *gin.Context) {
	appId := c.Param("id")
	courseType := services.CourseType{}
	err := c.ShouldBindJSON(&courseType)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.DeleteCourseType(appId, &courseType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "delete course type successfully"})
}
func (h *AppSettingHandler) GetClinics(c *gin.Context) {
	appId := c.Param("id")
	if appId == "" {
		c.JSON(400, gin.H{"error": "Invalid app id"})
		return
	}
	clinics, err := h.appSettingServ.ClinicSettingList(appId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": clinics})
}
func (h *AppSettingHandler) PostAddClinic(c *gin.Context) {
	appId := c.Param("id")
	if appId == "" {
		c.JSON(400, gin.H{"error": "Invalid app id"})
		return
	}
	clinicImpl := services.ClinicSettingImpl{}
	err := c.ShouldBindJSON(&clinicImpl)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.AddClinicSetting(appId, &clinicImpl)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "add clinic successfully"})

}
func (h *AppSettingHandler) PutUpdateClinic(c *gin.Context) {
	appId := c.Param("id")
	if appId == "" {
		c.JSON(400, gin.H{"error": "Invalid app id"})
		return
	}
	clinicImpl := services.ClinicSettingImpl{}
	err := c.ShouldBindJSON(&clinicImpl)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.UpdateClinicSetting(appId, &clinicImpl)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "update clinic successfully"})

}
func (h *AppSettingHandler) DeleteClinic(c *gin.Context) {
	appId := c.Param("id")
	if appId == "" {
		c.JSON(400, gin.H{"error": "Invalid app id"})
		return
	}
	clinicImpl := services.ClinicSettingImpl{}
	err := c.ShouldBindJSON(&clinicImpl)
	cusErr := utils.NewCustomErrorHandler(c)
	if err != nil {
		cusErr.ValidateError(err)
		return
	}
	err = h.appSettingServ.DeleteClinicSetting(appId, &clinicImpl)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "delete clinic successfully"})

}
