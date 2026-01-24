package models

import (
	"errors"
	"time"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailNotFound      = errors.New("email does not exist")
	ErrInvalidPassword    = errors.New("invalid password")
)

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	SQLModel
	FullName     string     `json:"full_name" gorm:"full_name"`
	HashPassword string     `json:"hash_password" gorm:"hash_password"`
	DateOfBirth  *time.Time `json:"date_of_birth,omitempty" gorm:"date_of_birth"`
	Gender       *string    `json:"gender,omitempty" gorm:"gender"`
	IdCardNumber *string    `json:"id_card_number" gorm:"id_card_number"`
	Email        string     `json:"email" gorm:"email"`
	PhoneNumber  *string    `json:"phone_number" gorm:"phone_number"`
	Avatar       *string    `json:"avatar,omitempty" gorm:"avatar"`
}

func (User) TableName() string {
	return "users"
}
