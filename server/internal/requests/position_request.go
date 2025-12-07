package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vlahanam/company-management/common"
)

type CreatePositionRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Level       *int    `json:"level,omitempty"`
}

type UpdatePositionRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Level       *int    `json:"level,omitempty"`
}

type ListPositionRequest struct {
	common.Paging
	CompanyID *uint64 `json:"company_id,omitempty"`
}

func (r CreatePositionRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.RuneLength(1, 200)),
	)
}

func (r UpdatePositionRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.When(r.Name != nil, validation.RuneLength(1, 200))),
	)
}
