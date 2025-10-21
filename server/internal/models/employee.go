package models

import (
	"errors"
	"time"
)

var (
	ErrEmployeeNotFound   = errors.New("employee not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type Employee struct {
	ID           int64      `gorm:"id" json:"-"`
	FullName     string     `gorm:"id" json:"full_name"`
	HashPassword string     `gorm:"id" json:"hash_password"`
	DateOfBirth  *time.Time `gorm:"id" json:"date_of_birth,omitempty"`
	Gender       *string    `gorm:"id" json:"gender,omitempty"`
	IdCardNumber *string    `gorm:"id" json:"id_card_number"`
	Email        string     `gorm:"id" json:"email"`
	PhoneNumber  *string    `gorm:"id" json:"phone_number"`
	Avatar       *string    `gorm:"id" json:"avatar,omitempty"`
	CreatedAt    *time.Time `gorm:"id" json:"created_at,omitempty"`
	UpdatedAt    *time.Time `gorm:"id" json:"updated_at,omitempty"`
}

func (Employee) TableName() string {
	return "employees"
}
