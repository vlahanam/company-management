package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vlahanam/company-management/common"
)

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateRoleRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ListRoleRequest struct {
	common.Paging
}

func (r CreateRoleRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.RuneLength(1, 100)),
	)
}

func (r UpdateRoleRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.When(r.Name != nil, validation.RuneLength(1, 100))),
	)
}
