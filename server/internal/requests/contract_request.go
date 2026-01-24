package requests

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vlahanam/company-management/common"
)

type CreateContractRequest struct {
	UserID         uint64  `json:"user_id"`
	CompanyID      uint64  `json:"company_id"`
	PositionID     *uint64 `json:"position_id,omitempty"`
	ContractNumber string  `json:"contract_number"`
	ContractType   string  `json:"contract_type"`
	StartDate      string  `json:"start_date"` // Format: "2006-01-02"
	EndDate        *string `json:"end_date,omitempty"`
	Salary         float64 `json:"salary"`
	Status         string  `json:"status"`
	FilePath       *string `json:"file_path,omitempty"`
	Notes          *string `json:"notes,omitempty"`
}

type UpdateContractRequest struct {
	ContractType *string    `json:"contract_type,omitempty"`
	StartDate    *string    `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	Salary       *float64   `json:"salary,omitempty"`
	Status       *string    `json:"status,omitempty"`
	FilePath     *string    `json:"file_path,omitempty"`
	Notes        *string    `json:"notes,omitempty"`
}

type ListContractRequest struct {
	common.Paging
	UserID    *uint64 `json:"user_id,omitempty"`
	CompanyID *uint64 `json:"company_id,omitempty"`
	Status    *string `json:"status,omitempty"`
	Type      *string `json:"type,omitempty"`
}

func (r CreateContractRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required),
		validation.Field(&r.CompanyID, validation.Required),
		validation.Field(&r.ContractNumber, validation.Required, validation.RuneLength(1, 50)),
		validation.Field(&r.ContractType, validation.Required, validation.In("Probation", "Fixed-term", "Permanent", "Freelance", "Internship")),
		validation.Field(&r.StartDate, validation.Required),
		validation.Field(&r.Salary, validation.Required, validation.Min(0.0)),
		validation.Field(&r.Status, validation.Required, validation.In("Active", "Pending", "Expired", "Terminated")),
	)
}

func (r UpdateContractRequest) Validation() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ContractType, validation.When(r.ContractType != nil, validation.In("Probation", "Fixed-term", "Permanent", "Freelance", "Internship"))),
		validation.Field(&r.Status, validation.When(r.Status != nil, validation.In("Active", "Pending", "Expired", "Terminated"))),
		validation.Field(&r.Salary, validation.When(r.Salary != nil, validation.Min(0.0))),
	)
}
