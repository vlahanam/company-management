package requests

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vlahanam/company-management/common"
)

type ListUserRequest struct {
	common.Paging
	KeyWord    *string `json:"keyword,omitempty"`
	CompanyID  *int64  `json:"company_id,omitempty"`
	PositionID *int64  `json:"position_id,omitempty"`
}

type UpdateUserRequest struct {
	FullName     *string    `json:"full_name,omitempty"`
	Email        *string    `json:"email,omitempty"`
	DateOfBirth  *time.Time `json:"date_of_birth,omitempty"`
	Gender       *string    `json:"gender,omitempty"`
	IdCardNumber *string    `json:"id_card_number,omitempty"`
	PhoneNumber  *string    `json:"phone_number,omitempty"`
	Avatar       *string    `json:"avatar,omitempty"`
}

func (r UpdateUserRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.When(r.Email != nil, isValidEmail())),
		validation.Field(&r.FullName, validation.When(r.FullName != nil, validation.RuneLength(1, 100))),
		validation.Field(&r.Gender, validation.When(r.Gender != nil, validation.In("Male", "Female", "Other"))),
		validation.Field(&r.IdCardNumber, validation.When(r.IdCardNumber != nil, validation.RuneLength(9, 12))),
	)
}


