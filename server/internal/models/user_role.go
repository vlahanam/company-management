package models

import "time"

type UserRole struct {
	UserID     int64      `json:"user_id" gorm:"column:user_id"`
	RoleID     int64      `json:"role_id" gorm:"column:role_id"`
	AssignedAt *time.Time `json:"assigned_at,omitempty" gorm:"column:assigned_at;"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
