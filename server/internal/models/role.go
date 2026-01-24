package models

import "time"

type Role struct {
	ID          int64      `json:"id" gorm:"column:id"`
	Name        string     `json:"name" gorm:"column:name"`
	Description string     `json:"description,omitempty" gorm:"column:description"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
}

func (Role) TableName() string {
	return "roles"
}

func (r *Role) GetRoleName() string {
	if name, ok := RoleNames[r.ID]; ok {
		return name
	}
	return "Unknown Role"
}

const (
	RoleSuperAdmin int64 = iota + 1
	RoleAdmin
	RoleHRManager
	RoleHRStaff
	RoleFinanceMgr
	RoleAccountant
	RoleSalesManager
	RoleSalesStaff
	RoleProductMgr
	RoleEmployee
)

var RoleNames = map[int64]string{
	RoleSuperAdmin:   "Super Admin",
	RoleAdmin:        "Admin",
	RoleHRManager:    "HR Manager",
	RoleHRStaff:      "HR Staff",
	RoleFinanceMgr:   "Finance Manager",
	RoleAccountant:   "Accountant",
	RoleSalesManager: "Sales Manager",
	RoleSalesStaff:   "Sales Staff",
	RoleProductMgr:   "Product Manager",
	RoleEmployee:     "Employee",
}
