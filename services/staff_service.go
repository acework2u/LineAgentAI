package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"linechat/repository"
	"linechat/utils"
	"time"
)

type staffServiceImpl struct {
	staffRep repository.StaffRepository
}

func NewStaffService(staffRepo repository.StaffRepository) StaffService {
	return &staffServiceImpl{staffRep: staffRepo}
}
func (s *staffServiceImpl) Register(staff *StaffRegister) error {
	if staff.Password == "" {
		return errors.New("password is empty")
	}
	if staff.Email == "" {
		return errors.New("email is empty")
	}
	if staff.Name == "" {
		return errors.New("name is empty")
	}

	password, err := utils.HashPassword(staff.Password)

	if err != nil {
		return err
	}
	staff.Password = password

	res := s.staffRep.CreateStaff(&repository.Staff{
		Name:       staff.Name,
		Email:      staff.Email,
		Password:   staff.Password,
		Tel:        staff.Tel,
		Role:       staff.Role,
		ClinicId:   0,
		ClinicName: "",
		Status:     true,
		LastLogin:  0,
		LineId:     staff.LineId,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if res != nil {
		return res
	}

	return nil
}
func (s *staffServiceImpl) LoginStaff(staff *StaffLogin) (string, error) {
	if staff.Email == "" {
		return "", errors.New("email is empty")
	}
	if staff.Password == "" {
		return "", errors.New("password is empty")
	}

	res, err := s.staffRep.GetStaffByEmail(staff.Email)
	if err != nil {
		return "", err
	}
	// check password
	if err := utils.ComparePassword(staff.Password, res.Password); err != nil {
		return "", errors.New("password is incorrect")
	}
	// Generate token
	token, err := utils.GenerateToken(time.Hour*24, res.LineId, res)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (s *staffServiceImpl) GetStaffById(id string) (*StaffResponse, error) {

	if id == "" {
		return nil, errors.New("staff id is empty")
	}
	res, err := s.staffRep.GetStaffById(id)
	if err != nil {
		return nil, err
	}
	staff := &StaffResponse{
		Id:         res.Id.String(),
		Name:       res.Name,
		Email:      res.Email,
		Role:       res.Role,
		ClinicId:   res.ClinicId,
		ClinicName: res.ClinicName,
		Status:     res.Status,
		LastLogin:  res.LastLogin,
		LineId:     res.LineId,
		UpdatedAt:  res.UpdatedAt,
	}

	return staff, nil
}
func (s *staffServiceImpl) GetStaffs() ([]StaffResponse, error) {

	res, err := s.staffRep.GetStaffs()
	if err != nil {
		return nil, err
	}
	staffs := []StaffResponse{}
	for _, staff := range res {
		staffs = append(staffs, StaffResponse{
			Id:         staff.Id.Hex(),
			Name:       staff.Name,
			Email:      staff.Email,
			Role:       staff.Role,
			ClinicId:   staff.ClinicId,
			ClinicName: staff.ClinicName,
			Status:     staff.Status,
			LastLogin:  staff.LastLogin,
			LineId:     staff.LineId,
			UpdatedAt:  staff.UpdatedAt,
		})
	}
	return staffs, nil
}
func (s *staffServiceImpl) UpdateStaff(staff *Staff) error {

	if staff.Id == "" {
		return errors.New("staff id is empty")
	}
	staffId, err := primitive.ObjectIDFromHex(staff.Id)
	if err != nil {
		return err
	}
	err = s.staffRep.UpdateStaff(&repository.Staff{
		Id:         staffId,
		Name:       staff.Name,
		Email:      staff.Email,
		Tel:        staff.Tel,
		Role:       staff.Role,
		ClinicId:   0,
		ClinicName: staff.ClinicName,
		Status:     staff.Status,
		LastLogin:  staff.LastLogin,
		LineId:     staff.LineId,
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
func (s staffServiceImpl) DeleteStaff(staff *Staff) error {
	if staff.Email == "" {
		return errors.New("staff id is empty")
	}
	err := s.staffRep.DeleteStaff(&repository.Staff{
		Email: staff.Email,
	})
	if err != nil {
		return err
	}

	return nil
}
