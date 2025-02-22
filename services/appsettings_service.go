package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"linechat/repository"
	"time"
)

type AppSettingsServiceImpl struct {
	appSettingRepo repository.AppSettingsRepository
}

func NewAppSettingsService(appSettingRepo repository.AppSettingsRepository) AppSettingsService {
	return &AppSettingsServiceImpl{
		appSettingRepo: appSettingRepo,
	}
}
func (s *AppSettingsServiceImpl) CreateAppSettings(settings *AppSettings) error {

	if settings.Title == "" {
		return errors.New("title is required")
	}

	convertMemberType := make([]*repository.MemberTypeSettingImpl, 0, len(settings.MemberType))
	for _, memberType := range settings.MemberType {
		convertMemberType = append(convertMemberType, &repository.MemberTypeSettingImpl{
			Title:  memberType.Title,
			Status: true,
		})
	}

	// the clinic add to app setting
	convertClinicSetting := make([]*repository.ClinicSettingImpl, 0, len(settings.ClinicSetting))
	for _, clinicSetting := range settings.ClinicSetting {
		convertClinicSetting = append(convertClinicSetting, &repository.ClinicSettingImpl{
			Title:  clinicSetting.Title,
			Status: true,
		})
	}

	err := s.appSettingRepo.CreateAppSettings(&repository.AppSettings{
		Id:            primitive.ObjectID{},
		Name:          settings.Title,
		MemberType:    convertMemberType,
		ClinicSetting: convertClinicSetting,
		Status:        true,
		CreatedAt:     time.Now().Local().Unix(),
		UpdatedAt:     time.Now().Local().Unix(),
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *AppSettingsServiceImpl) GetAppSettings() (*AppSettings, error) {

	res, err := s.appSettingRepo.GetAppSettings()
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}
	convertClinic := make([]*ClinicSettingImpl, 0, len(res.ClinicSetting))
	for _, clinic := range res.ClinicSetting {
		convertClinic = append(convertClinic, &ClinicSettingImpl{
			Title:  clinic.Title,
			Status: clinic.Status,
		})
	}
	convertMemberType := make([]*MemberTypeImpl, 0, len(res.MemberType))
	for _, memberType := range res.MemberType {
		convertMemberType = append(convertMemberType, &MemberTypeImpl{
			Title:  memberType.Title,
			Status: memberType.Status,
		})
	}

	appSettings := &AppSettings{
		Id:            res.Id.Hex(),
		Title:         res.Name,
		MemberType:    convertMemberType,
		ClinicSetting: convertClinic,
	}

	return appSettings, nil

}
func (s *AppSettingsServiceImpl) UpdateAppSettings(settings *AppSettings) error {
	return nil
}
func (s *AppSettingsServiceImpl) AddMemberType(appId string, memberType *MemberTypeImpl) error {
	return nil
}
