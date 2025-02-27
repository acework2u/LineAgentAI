package services

import (
	"errors"
	"github.com/google/uuid"
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
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
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
	courses := make([]*Course, 0, len(res.Courses))
	for _, course := range res.Courses {
		courses = append(courses, &Course{
			Name:   course.Name,
			Type:   course.Type,
			Desc:   course.Desc,
			Img:    course.Img,
			Status: course.Status,
		})
	}
	appSettings := &AppSettings{
		Id:            res.Id.Hex(),
		Title:         res.Name,
		MemberType:    convertMemberType,
		ClinicSetting: convertClinic,
		Courses:       courses,
	}

	return appSettings, nil

}
func (s *AppSettingsServiceImpl) UpdateAppSettings(settings *AppSettings) error {
	return nil
}
func (s *AppSettingsServiceImpl) AddMemberType(appId string, memberType *MemberTypeImpl) error {

	if appId == "" {
		return errors.New("app id is empty")
	}
	memberTypeId, _ := uuid.NewUUID()
	err := s.appSettingRepo.AddMemberType(appId, &repository.MemberTypeSettingImpl{
		Id:     memberTypeId.String(),
		Title:  memberType.Title,
		Status: true,
	})
	if err != nil {
		return err
	}

	return nil

}
func (s *AppSettingsServiceImpl) MemberTypesList(appId string) ([]*MemberTypeImpl, error) {
	if appId == "" {
		return nil, errors.New("app id is empty")
	}
	res, err := s.appSettingRepo.MemberTypesetting(appId)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
		return nil, err
	}

	convertMemberType := make([]*MemberTypeImpl, 0, len(res))
	for _, memberType := range res {
		convertMemberType = append(convertMemberType, &MemberTypeImpl{
			Id:     memberType.Id,
			Title:  memberType.Title,
			Status: memberType.Status,
		})
	}
	return convertMemberType, nil
}
func (s *AppSettingsServiceImpl) UpdateMemberType(appId string, memberType *MemberTypeImpl) error {
	if appId == "" {
		return errors.New("app id is empty")
	}
	err := s.appSettingRepo.UpdateMemberType(appId, &repository.MemberTypeSettingImpl{
		Id:     memberType.Id,
		Title:  memberType.Title,
		Status: memberType.Status,
	})
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("member type not found")
		}
		return err
	}
	return nil
}
func (s *AppSettingsServiceImpl) DeleteMemberType(appId string, memberType *MemberTypeImpl) error {
	if appId == "" {
		return errors.New("app id is empty")
	}
	err := s.appSettingRepo.DeleteMemberType(appId, &repository.MemberTypeSettingImpl{
		Title:  memberType.Title,
		Status: memberType.Status,
	})
	if err != nil {
		return err
	}
	return nil

}
func (s *AppSettingsServiceImpl) DeleteAppSettings(appId string) error {
	if appId == "" {
		return errors.New("app id is empty")
	}
	// delete app setting
	err := s.appSettingRepo.DeleteAppSettings(appId)
	if err != nil {
		return err
	}
	return nil
}
func (s *AppSettingsServiceImpl) AddCourse(appId string, course *Course) error {

	if appId == "" {
		return errors.New("app id is empty")
	}
	err := s.appSettingRepo.AddCourse(appId, &repository.Course{
		Name:   course.Name,
		Type:   course.Type,
		Desc:   course.Desc,
		Img:    course.Img,
		Status: true,
	})
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil
		}
		return err
	}

	return nil
}
func (s *AppSettingsServiceImpl) UpdateCourse(appId string, course *Course) error {
	if appId == "" {
		return errors.New("app id is empty")
	}
	err := s.appSettingRepo.UpdateCourse(appId, &repository.Course{
		Id:     course.Id,
		Name:   course.Name,
		Type:   course.Type,
		Desc:   course.Desc,
		Img:    course.Img,
		Status: course.Status,
	})
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("course not found")
		}
		return err
	}

	return nil
}
func (s *AppSettingsServiceImpl) DeleteCourse(appId string, course *Course) error {
	if appId == "" {
		return errors.New("app id is empty")
	}
	err := s.appSettingRepo.DeleteCourse(appId, &repository.Course{
		Id:     course.Id,
		Name:   course.Name,
		Type:   course.Type,
		Desc:   course.Desc,
		Img:    course.Img,
		Status: course.Status,
	})
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("course not found")
		}
		return err
	}
	return nil
}
func (s *AppSettingsServiceImpl) CourseList(appId string) ([]*Course, error) {
	if appId == "" {
		return nil, errors.New("app id is empty")
	}

	courses, err := s.appSettingRepo.CourseListSetting(appId)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
		return nil, err
	}
	convertCourse := make([]*Course, 0, len(courses))
	for _, course := range courses {
		convertCourse = append(convertCourse, &Course{
			Id:     course.Id,
			Name:   course.Name,
			Type:   course.Type,
			Desc:   course.Desc,
			Img:    course.Img,
			Status: course.Status,
		})
	}
	return convertCourse, nil
}
func (s *AppSettingsServiceImpl) AddCourseType(appId string, courseType *CourseType) error {
	if appId == "" {
		return errors.New("app id is empty")
	}
	courseTypeId, _ := uuid.NewUUID()
	if courseType.Name == "" {
		return errors.New("course type name is empty")
	}
	err := s.appSettingRepo.AddCourseType(appId, &repository.CourseType{
		Id:   courseTypeId.String(),
		Name: courseType.Name,
	})
	if err != nil {
		return err
	}
	return nil

}
func (s *AppSettingsServiceImpl) CourseTypeList(appId string) ([]*CourseType, error) {
	if appId == "" {
		return nil, errors.New("app id is empty")
	}
	result, err := s.appSettingRepo.CourseTypeList(appId)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
	}
	convertCourseType := make([]*CourseType, 0, len(result))
	for _, courseType := range result {
		convertCourseType = append(convertCourseType, &CourseType{
			Id:   courseType.Id,
			Name: courseType.Name,
		})
	}
	return convertCourseType, nil
}
func (s *AppSettingsServiceImpl) UpdateCourseType(appId string, courseType *CourseType) error {
	if appId == "" {
		return errors.New("app id is empty")
	}
	err := s.appSettingRepo.UpdateCourseType(appId, &repository.CourseType{
		Id:   courseType.Id,
		Name: courseType.Name,
	})
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("course type not found")
		}
		return err
	}
	return nil

}
func (s *AppSettingsServiceImpl) DeleteCourseType(appId string, courseType *CourseType) error {
	if appId == "" {
		return errors.New("app id is empty")
	}
	err := s.appSettingRepo.DeleteCourseType(appId, &repository.CourseType{
		Id:   courseType.Id,
		Name: courseType.Name,
	})
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("course type not found")
		}
		return err
	}
	return nil
}
