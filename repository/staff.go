package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type StaffRepository interface {
	GetStaffs() ([]Staff, error)
	CreateStaff(staff *Staff) error
	UpdateStaff(staff *Staff) error
	DeleteStaff(staff *Staff) error
	GetStaffById(id string) (*Staff, error)
	GetStaffByEmail(email string) (*Staff, error)
}

type Staff struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Email      string             `bson:"email"`
	Password   string             `bson:"password"`
	Tel        string             `bson:"tel"`
	Role       string             `bson:"role"`
	ClinicId   int                `bson:"clinic_id"`
	ClinicName string             `bson:"clinic_name"`
	Status     bool               `bson:"status"`
	LastLogin  int64              `bson:"last_login"`
	LineId     string             `bson:"lineId"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

type StaffLogin struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
type StaffRegister struct {
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Tel       string    `bson:"tel"`
	Role      string    `bson:"role"`
	ClinicId  int       `bson:"clinic_id"`
	Status    bool      `bson:"status"`
	LastLogin int64     `bson:"last_login"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
