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
	AddCourse(appId string, course *Course) error
	UpdateCourse(appId string, course *Course) error
	DeleteCourse(appId string, course *Course) error
	CourseList(appId string) ([]*Course, error)
	AddCourseType(appId string, courseType *CourseType) error
	CourseTypeList(appId string) ([]*CourseType, error)
	UpdateCourseType(appId string, courseType *CourseType) error
	DeleteCourseType(appId string, courseType *CourseType) error
	AddClinicSetting(appId string, clinicSetting *ClinicSettingImpl) error
	ClinicSettingList(appId string) ([]*ClinicSettingImpl, error)
	UpdateClinicSetting(appId string, clinicSetting *ClinicSettingImpl) error
	DeleteClinicSetting(appId string, clinicSetting *ClinicSettingImpl) error
	AddBanner(appId string, banner *Banner) error
	UpdateBanner(appId string, banner *Banner) error
	DeleteBanner(appId string, bannerId string) error
	BannerList(appId string) ([]*Banner, error)
}

type AppSettings struct {
	Id            string               `json:"_id,omitempty"`
	Title         string               `json:"title" binding:"required"`
	MemberType    []*MemberTypeImpl    `json:"member_type"`
	ClinicSetting []*ClinicSettingImpl `json:"clinic_setting"`
	Courses       []*Course            `json:"courses"`
	CourseType    []*CourseType        `json:"course_type"`
	Banners       []*Banner            `json:"banners"`
	Status        bool                 `json:"status"`
	CreatedAt     int64                `json:"created_at,omitempty"`
	UpdatedAt     int64                `json:"updated_at,omitempty"`
}

type MemberTypeImpl struct {
	Id     string `json:"id,omitempty"`
	Title  string `json:"title" binding:"required"`
	Status bool   `json:"status"`
}

type ClinicSettingImpl struct {
	Id     string `json:"id"`
	Title  string `json:"title" binding:"required"`
	Status bool   `json:"status"`
}
type Course struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name" binding:"required"`
	Type   string `json:"type" binding:"required"`
	Desc   string `json:"desc"`
	Img    string `json:"img"`
	Status bool   `json:"status"`
}
type CourseType struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name" binding:"required"`
}
type Banner struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title" binding:"required"`
	Url   string `json:"url" binding:"required"`
}
