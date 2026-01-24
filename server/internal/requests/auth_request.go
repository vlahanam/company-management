package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (lr LoginRequest) Validation() error {
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Email, validation.Required, isValidEmail()),
		validation.Field(&lr.Password, validation.Required),
	)
}

func (rr RegisterRequest) Validation() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.FullName, validation.Required, validation.RuneLength(1, 100)),
		validation.Field(&rr.Email, validation.Required, isValidEmail()),
		validation.Field(&rr.Password, validation.Required),
	)
}

func (rr RefreshRequest) Validation() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.RefreshToken, validation.Required),
	)
}
