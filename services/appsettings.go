package services

type AppSettingsService interface {
	CreateAppSettings(settings *AppSettings) error
	GetAppSettings() (*AppSettings, error)
	UpdateAppSettings(settings *AppSettings) error
	AddMemberType(appId string, memberType *MemberTypeImpl) error
}

type AppSettings struct {
	Id            string               `json:"_id,omitempty"`
	Title         string               `json:"title" binding:"required"`
	MemberType    []*MemberTypeImpl    `json:"member_type"`
	ClinicSetting []*ClinicSettingImpl `json:"clinic_setting"`
}

type MemberTypeImpl struct {
	Title  string `json:"title" binding:"required"`
	Status bool   `json:"status"`
}

type ClinicSettingImpl struct {
	ClinicId int    `json:"clinic_id"`
	Title    string `json:"title" binding:"required"`
	Status   bool   `json:"status"`
}
