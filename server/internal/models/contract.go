package models

import (
	"errors"
	"time"
)

var (
	ErrContractNotFound = errors.New("contract not found")
)

// ContractType represents the type of employment contract
type ContractType string

const (
	ContractTypeProbation  ContractType = "Probation"
	ContractTypeFixedTerm  ContractType = "Fixed-term"
	ContractTypePermanent  ContractType = "Permanent"
	ContractTypeFreelance  ContractType = "Freelance"
	ContractTypeInternship ContractType = "Internship"
)

// ContractStatus represents the current status of a contract
type ContractStatus string

const (
	ContractStatusActive     ContractStatus = "Active"
	ContractStatusExpired    ContractStatus = "Expired"
	ContractStatusTerminated ContractStatus = "Terminated"
	ContractStatusPending    ContractStatus = "Pending"
)

type Contract struct {
	SQLModel
	UserID         uint64         `json:"user_id" gorm:"column:user_id"`
	CompanyID      uint64         `json:"company_id" gorm:"column:company_id"`
	PositionID     *uint64        `json:"position_id,omitempty" gorm:"column:position_id"`
	ContractNumber string         `json:"contract_number" gorm:"column:contract_number"`
	ContractType   ContractType   `json:"contract_type" gorm:"column:contract_type;default:'Fixed-term'"`
	StartDate      time.Time      `json:"start_date" gorm:"column:start_date"`
	EndDate        *time.Time     `json:"end_date,omitempty" gorm:"column:end_date"`
	Salary         float64        `json:"salary" gorm:"column:salary;type:decimal(15,2)"`
	Status         ContractStatus `json:"status" gorm:"column:status;default:'Pending'"`
	FilePath       *string        `json:"file_path,omitempty" gorm:"column:file_path"`
	Notes          *string        `json:"notes,omitempty" gorm:"column:notes"`

	// Relationships
	User     *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Company  *Company  `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Position *Position `json:"position,omitempty" gorm:"foreignKey:PositionID"`
}

func (Contract) TableName() string {
	return "contracts"
}
