package models

import (
	"errors"
	"time"
)

var (
	ErrCompanyNotFound = errors.New("company not found")
)

type Company struct {
	SQLModel
	Name        string     `json:"name" gorm:"column:name"`
	ParentID    *uint64    `json:"parent_id,omitempty" gorm:"column:parent_id"`
	Description *string    `json:"description,omitempty" gorm:"column:description"`
	FoundedDate *time.Time `json:"founded_date,omitempty" gorm:"column:founded_date"`
	Address     *string    `json:"address,omitempty" gorm:"column:address"`
	PhoneNumber *string    `json:"phone_number,omitempty" gorm:"column:phone_number"`
	Email       *string    `json:"email,omitempty" gorm:"column:email"`

	// Relationships
	Parent    *Company   `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children  []*Company `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Positions []Position `json:"positions,omitempty" gorm:"foreignKey:CompanyID"`
}

func (Company) TableName() string {
	return "companies"
}
