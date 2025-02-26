package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppSettingsRepository interface {
	CreateAppSettings(settings *AppSettings) error
	GetAppSettings() (*AppSettings, error)
	UpdateAppSettings(settings *AppSettings) error
	DeleteAppSettings(appId string) error
	AddMemberType(appId string, memberType *MemberTypeSettingImpl) error
	UpdateMemberType(appId string, memberType *MemberTypeSettingImpl) error
	DeleteMemberType(appId string, memberType *MemberTypeSettingImpl) error
	MemberTypesetting(appId string) ([]*MemberTypeSettingImpl, error)
	AddClinicSetting(appId string, clinicSetting *ClinicSettingImpl) error
	AddCourse(appId string, course *Course) error
	AddCourseType(appId string, courseType string) error
}

type AppSettings struct {
	Id            primitive.ObjectID       `bson:"_id,omitempty"`
	Name          string                   `bson:"name"`
	MemberType    []*MemberTypeSettingImpl `bson:"members_type,omitempty"`
	ClinicSetting []*ClinicSettingImpl     `bson:"clinic_setting,omitempty"`
	Courses       []*Course                `bson:"course_setting,omitempty"`
	Status        bool                     `bson:"status"`
	CreatedAt     int64                    `bson:"created_at,omitempty"`
	UpdatedAt     int64                    `bson:"updated_at,omitempty"`
}
type MemberTypeSettingImpl struct {
	Title  string `bson:"title"`
	Status bool   `bson:"status"`
}

type ClinicSettingImpl struct {
	ClinicId int    `bson:"clinic_id"`
	Title    string `bson:"title"`
	Status   bool   `bson:"status"`
}
type Course struct {
	Id     string `bson:"_id,omitempty"`
	Name   string `bson:"name"`
	Type   string `bson:"type"`
	Desc   string `bson:"desc"`
	Img    string `bson:"img"`
	Status bool   `bson:"status"`
}
