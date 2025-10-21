package models

import "time"

type Permission struct {
	ID          int64      `json:"id" gorm:"column:id"`
	Name        string     `json:"name" gorm:"column:name"`
	Description string     `json:"description,omitempty" gorm:"column:description"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
}

func (Permission) TableName() string {
	return "permissions"
}

func (p *Permission) GetPermissionName() string {
	if name, ok := PermissionNames[p.ID]; ok {
		return name
	}
	return "Unknown Permission"
}

const (
	PermissionCreateUser int64 = iota + 1
	PermissionReadUser
	PermissionUpdateUser
	PermissionDeleteUser
	PermissionManageRoles
	PermissionViewOwnReport
	PermissionViewSubordinateReport
	PermissionViewAllReports
	PermissionCreateReport
	PermissionDeleteReport
	PermissionManageFinances
	PermissionManageInventory
	PermissionPosition
	PermissionUpdatePosition
	PermissionDeletePosition
	PermissionCreateDepartment
	PermissionUpdateDepartment
	PermissionDeleteDepartment
	PermissionCreateCompany
	PermissionUpdateCompany
	PermissionDeleteCompany
	PermissionCreateContract
	PermissionUpdateContract
	PermissionDeleteContract
	PermissionApproveRequests
	PermissionCreateRequest
	PermissionUpdateRequest
	PermissionDeleteRequest
)

var PermissionNames = map[int64]string{
	PermissionCreateUser:            "Create User",
	PermissionReadUser:              "Read User",
	PermissionUpdateUser:            "Update User",
	PermissionDeleteUser:            "Delete User",
	PermissionManageRoles:           "Manage Roles",
	PermissionViewOwnReport:         "View Own Report",
	PermissionViewSubordinateReport: "View Subordinate Report",
	PermissionViewAllReports:        "View All Reports",
	PermissionCreateReport:          "Create Report",
	PermissionDeleteReport:          "Delete Report",
	PermissionManageFinances:        "Manage Finances",
	PermissionManageInventory:       "Manage Inventory",
	PermissionPosition:              "Create Position",
	PermissionUpdatePosition:        "Update Position",
	PermissionDeletePosition:        "Delete Position",
	PermissionCreateDepartment:      "Create Department",
	PermissionUpdateDepartment:      "Update Department",
	PermissionDeleteDepartment:      "Delete Department",
	PermissionCreateCompany:         "Create Company",
	PermissionUpdateCompany:         "Update Company",
	PermissionDeleteCompany:         "Delete Company",
	PermissionCreateContract:        "Create Contract",
	PermissionUpdateContract:        "Update Contract",
	PermissionDeleteContract:        "Delete Contract",
	PermissionApproveRequests:       "Approve Requests",
	PermissionCreateRequest:         "Create Request",
	PermissionUpdateRequest:         "Update Request",
	PermissionDeleteRequest:         "Delete Request",
}
