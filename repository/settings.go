package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppSettingsRepository interface {
	CreateAppSettings(settings *AppSettings) error
	GetAppSettings() (*AppSettings, error)
	UpdateAppSettings(settings *AppSettings) error
	AddMemberType(appId string, memberType *MemberTypeSettingImpl) error
}

type AppSettings struct {
	Id            primitive.ObjectID       `bson:"_id,omitempty"`
	Name          string                   `bson:"name"`
	MemberType    []*MemberTypeSettingImpl `bson:"members_type,omitempty"`
	ClinicSetting []*ClinicSettingImpl     `bson:"clinics_setting,omitempty"`
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
