package models

import "errors"

var (
	ErrPositionNotFound = errors.New("position not found")
)

type Position struct {
	SQLModel
	CompanyID   uint64  `json:"company_id" gorm:"column:company_id"`
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description,omitempty" gorm:"column:description"`
	Level       *int    `json:"level,omitempty" gorm:"column:level"`

	// Relationships
	Company *Company `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
}

func (Position) TableName() string {
	return "positions"
}
