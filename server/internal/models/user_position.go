package models

import (
	"errors"
	"time"
)

var (
	ErrUserPositionNotFound = errors.New("user position not found")
)

type UserPosition struct {
	SQLModel
	UserID     uint64     `json:"user_id" gorm:"column:user_id"`
	PositionID uint64     `json:"position_id" gorm:"column:position_id"`
	StartDate  time.Time  `json:"start_date" gorm:"column:start_date"`
	EndDate    *time.Time `json:"end_date,omitempty" gorm:"column:end_date"`
	IsPrimary  bool       `json:"is_primary" gorm:"column:is_primary;default:false"`
	
	// Relationships
	User     *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Position *Position `json:"position,omitempty" gorm:"foreignKey:PositionID"`
}

func (UserPosition) TableName() string {
	return "user_positions"
}

