package services

type AppSettingsService interface {
	CreateAppSettings(settings *AppSettings) error
	GetAppSettings() (*AppSettings, error)
	UpdateAppSettings(settings *AppSettings) error
	DeleteAppSettings(appId string) error
	AddMemberType(appId string, memberType *MemberTypeImpl) error
	UpdateMemberType(appId string, memberType *MemberTypeImpl) error
	DeleteMemberType(appId string, memberType *MemberTypeImpl) error
	MemberTypesList(appId string) ([]*MemberTypeImpl, error)
}

type AppSettings struct {
	Id            string               `json:"_id,omitempty"`
	Title         string               `json:"title" binding:"required"`
	MemberType    []*MemberTypeImpl    `json:"member_type"`
	ClinicSetting []*ClinicSettingImpl `json:"clinic_setting"`
	Courses       []*Course            `json:"courses"`
	Status        bool                 `json:"status"`
	CreatedAt     int64                `json:"created_at,omitempty"`
	UpdatedAt     int64                `json:"updated_at,omitempty"`
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
type Course struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Desc   string `json:"desc"`
	Img    string `json:"img"`
	Status bool   `json:"status"`
}
