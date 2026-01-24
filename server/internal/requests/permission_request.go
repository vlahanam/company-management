package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vlahanam/company-management/common"
)

type CreatePermissionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdatePermissionRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ListPermissionRequest struct {
	common.Paging
}

func (r CreatePermissionRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.RuneLength(1, 100)),
	)
}

func (r UpdatePermissionRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.When(r.Name != nil, validation.RuneLength(1, 100))),
	)
}
