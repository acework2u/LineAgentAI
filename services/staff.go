package services

import "time"

type StaffService interface {
	Register(staff *StaffRegister) error
	LoginStaff(staff *StaffLogin) (string, error)
	GetStaffById(id string) (*StaffResponse, error)
	GetStaffs() ([]StaffResponse, error)
	UpdateStaff(staff *Staff) error
	DeleteStaff(staff *Staff) error
}

type Staff struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Tel        string    `json:"tel"`
	Role       string    `json:"role"`
	ClinicId   string    `json:"clinic_id"`
	ClinicName string    `json:"clinic_name"`
	Status     bool      `json:"status"`
	LastLogin  int64     `json:"last_login"`
	LineId     string    `json:"lineId"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type StaffResponse struct {
	Id         string
	Name       string
	Email      string
	Tel        string
	Role       string
	ClinicId   int
	ClinicName string
	Status     bool
	LastLogin  int64
	LineId     string
	UpdatedAt  time.Time
}
type StaffUpdateImpl struct {
	Id         string
	Name       string
	Email      string
	Tel        string
	Role       string
	ClinicId   int
	ClinicName string
	Status     bool
	LastLogin  int64
	LineId     string
	UpdatedAt  time.Time
}
type StaffLogin struct {
	Email    string
	Password string
}
type StaffRegister struct {
	Name      string `json:"name" binding:"required" validate:"required,min=3,max=20"`
	Email     string `json:"email" binding:"required" validate:"required,email"`
	Password  string `json:"password" binding:"required" validate:"required,min=6"`
	Role      string
	ClinicId  string
	Status    bool
	LastLogin int64
	CreatedAt time.Time
	UpdatedAt time.Time
	LineId    string
	Tel       string
}
