package models

import "time"

type RolePermission struct {
	RoleID       int64      `json:"role_id" gorm:"column:role_id"`
	PermissionID int64      `json:"permission_id" gorm:"column:permission_id"`
	GrantedAt    *time.Time `json:"granted_at,omitempty" gorm:"column:granted_at;"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

var RolePermissions = map[int64][]int64{
	// Super Admin
	RoleSuperAdmin: {
		PermissionCreateUser,
		PermissionReadUser,
		PermissionUpdateUser,
		PermissionDeleteUser,
		PermissionManageRoles,
		PermissionViewOwnReport,
		PermissionViewSubordinateReport,
		PermissionViewAllReports,
		PermissionCreateReport,
		PermissionDeleteReport,
		PermissionManageFinances,
		PermissionManageInventory,
		PermissionPosition,
		PermissionUpdatePosition,
		PermissionDeletePosition,
		PermissionCreateDepartment,
		PermissionUpdateDepartment,
		PermissionDeleteDepartment,
		PermissionCreateCompany,
		PermissionUpdateCompany,
		PermissionDeleteCompany,
		PermissionCreateContract,
		PermissionUpdateContract,
		PermissionDeleteContract,
		PermissionApproveRequests,
		PermissionCreateRequest,
		PermissionUpdateRequest,
		PermissionDeleteRequest,
	},

	// Admin
	RoleAdmin: {
		PermissionCreateUser,
		PermissionReadUser,
		PermissionUpdateUser,
		PermissionDeleteUser,
		PermissionManageRoles,
		PermissionViewOwnReport,
		PermissionCreateReport,
		PermissionDeleteReport,
		PermissionManageFinances,
		PermissionManageInventory,
		PermissionPosition,
		PermissionUpdatePosition,
		PermissionCreateDepartment,
		PermissionUpdateDepartment,
		PermissionCreateCompany,
		PermissionUpdateCompany,
		PermissionCreateContract,
		PermissionUpdateContract,
		PermissionApproveRequests,
		PermissionCreateRequest,
		PermissionUpdateRequest,
	},

	// HR Manager
	RoleHRManager: {
		PermissionCreateUser,
		PermissionReadUser,
		PermissionUpdateUser,
		PermissionDeleteUser,
		PermissionManageRoles,
		PermissionViewOwnReport,
		PermissionCreateReport,
		PermissionDeleteReport,
		PermissionPosition,
		PermissionUpdatePosition,
		PermissionCreateDepartment,
		PermissionUpdateDepartment,
		PermissionDeleteDepartment,
		PermissionCreateCompany,
		PermissionUpdateCompany,
		PermissionCreateContract,
		PermissionUpdateContract,
		PermissionApproveRequests,
		PermissionCreateRequest,
		PermissionUpdateRequest,
	},

	// HR Staff
	RoleHRStaff: {
		PermissionReadUser,
		PermissionUpdateUser,
		PermissionViewOwnReport,
		PermissionCreateReport,
		PermissionDeleteReport,
		PermissionPosition,
		PermissionCreateDepartment,
		PermissionUpdateDepartment,
		PermissionDeleteDepartment,
		PermissionCreateRequest,
		PermissionUpdateRequest,
	},

	// Finance Manager
	RoleFinanceMgr: {
		PermissionCreateUser,
		PermissionReadUser,
		PermissionUpdateUser,
		PermissionDeleteUser,
		PermissionViewOwnReport,
		PermissionViewSubordinateReport,
		PermissionViewAllReports,
		PermissionManageFinances,
		PermissionManageInventory,
		PermissionCreateReport,
		PermissionDeleteReport,
		PermissionApproveRequests,
		PermissionCreateRequest,
		PermissionUpdateRequest,
	},

	// Accountant
	RoleAccountant: {
		PermissionReadUser,
		PermissionViewOwnReport,
		PermissionViewSubordinateReport,
		PermissionViewAllReports,
		PermissionManageFinances,
		PermissionCreateReport,
		PermissionDeleteReport,
		PermissionCreateRequest,
		PermissionUpdateRequest,
	},

	// Sales Manager
	RoleSalesManager: {
		PermissionCreateUser,
		PermissionReadUser,
		PermissionUpdateUser,
		PermissionDeleteUser,
		PermissionManageRoles,
		PermissionViewOwnReport,
		PermissionViewSubordinateReport,
		PermissionViewAllReports,
		PermissionCreateReport,
		PermissionDeleteReport,
		PermissionApproveRequests,
		PermissionCreateRequest,
		PermissionUpdateRequest,
	},

	// Sales Staff
	RoleSalesStaff: {
		PermissionReadUser,
		PermissionViewOwnReport,
		PermissionViewSubordinateReport,
		PermissionCreateReport,
		PermissionDeleteReport,
		PermissionCreateRequest,
		PermissionUpdateRequest,
	},

	// Product Manager
	RoleProductMgr: {
		PermissionCreateUser,
		PermissionReadUser,
		PermissionUpdateUser,
		PermissionDeleteUser,
		PermissionManageRoles,
		PermissionViewOwnReport,
		PermissionViewSubordinateReport,
		PermissionCreateReport,
		PermissionDeleteReport,
		PermissionCreateDepartment,
		PermissionUpdateDepartment,
		PermissionCreateContract,
		PermissionUpdateContract,
		PermissionCreateRequest,
		PermissionUpdateRequest,
	},

	// Employee
	RoleEmployee: {
		PermissionReadUser,
		PermissionViewOwnReport,
		PermissionCreateRequest,
		PermissionUpdateRequest,
	},
}
