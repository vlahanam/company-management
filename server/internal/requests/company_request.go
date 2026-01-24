package requests

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vlahanam/company-management/common"
)

type CreateCompanyRequest struct {
	Name        string     `json:"name"`
	ParentID    *uint64    `json:"parent_id,omitempty"`
	Description *string    `json:"description,omitempty"`
	FoundedDate *time.Time `json:"founded_date,omitempty"`
	Address     *string    `json:"address,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Email       *string    `json:"email,omitempty"`
}

type UpdateCompanyRequest struct {
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	FoundedDate *time.Time `json:"founded_date,omitempty"`
	Address     *string    `json:"address,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Email       *string    `json:"email,omitempty"`
}

type ListCompanyRequest struct {
	common.Paging
	Keyword  *string `json:"keyword,omitempty"`
	ParentID *uint64 `json:"parent_id,omitempty"`
}

func (r CreateCompanyRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.RuneLength(1, 200)),
		validation.Field(&r.Email, validation.When(r.Email != nil, isValidEmail())),
	)
}

func (r UpdateCompanyRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.When(r.Name != nil, validation.RuneLength(1, 200))),
		validation.Field(&r.Email, validation.When(r.Email != nil, isValidEmail())),
	)
}
