package main

import (
	"fmt"
	"log"
	"time"

	"github.com/vlahanam/company-management/internal/initialize"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/utils"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting database seeding...")

	// Load config and connect to database
	cfg := initialize.LoadConfig()
	db := initialize.InitMysql(cfg)

	// Seed roles
	if err := seedRoles(db); err != nil {
		log.Fatalf("Failed to seed roles: %v", err)
	}

	// Seed permissions
	if err := seedPermissions(db); err != nil {
		log.Fatalf("Failed to seed permissions: %v", err)
	}

	// Seed role_permissions
	if err := seedRolePermissions(db); err != nil {
		log.Fatalf("Failed to seed role permissions: %v", err)
	}

	// Seed users
	if err := seedUser(db); err != nil {
		log.Fatalf("Failed to seed user: %v", err)
	}

	// Seed user_roles
	if err := seedUserRole(db); err != nil {
		log.Fatalf("Failed to seed user role: %v", err)
	}

	fmt.Println("Database seeding completed!")
}

func seedRoles(db *gorm.DB) error {
	now := time.Now()

	// Build roles array from RoleNames map
	roles := make([]models.Role, 0, len(models.RoleNames))
	for roleID, roleName := range models.RoleNames {
		roles = append(roles, models.Role{
			ID:          roleID,
			Name:        roleName,
			Description: getRoleDescription(roleID),
			CreatedAt:   &now,
		})
	}

	// Batch insert all roles
	if err := db.Create(&roles).Error; err != nil {
		return fmt.Errorf("failed to create roles: %w", err)
	}

	return nil
}

// getRoleDescription returns description for each role
func getRoleDescription(roleID int64) string {
	descriptions := map[int64]string{
		models.RoleSuperAdmin:   "Full system access with all permissions",
		models.RoleAdmin:        "Administrative access to manage users and system",
		models.RoleHRManager:    "Manage human resources and employee data",
		models.RoleHRStaff:      "HR operations and employee management",
		models.RoleFinanceMgr:   "Manage financial operations and reports",
		models.RoleAccountant:   "Handle accounting and financial records",
		models.RoleSalesManager: "Manage sales team and operations",
		models.RoleSalesStaff:   "Sales operations and customer relations",
		models.RoleProductMgr:   "Manage products and development",
		models.RoleEmployee:     "Basic employee access",
	}
	return descriptions[roleID]
}

func seedPermissions(db *gorm.DB) error {
	now := time.Now()

	// Build permissions array from PermissionNames map
	permissions := make([]models.Permission, 0, len(models.PermissionNames))
	for permissionID, permissionName := range models.PermissionNames {
		permissions = append(permissions, models.Permission{
			ID:          permissionID,
			Name:        permissionName,
			Description: getPermissionDescription(permissionID),
			CreatedAt:   &now,
		})
	}

	// Batch insert all permissions
	if err := db.Create(&permissions).Error; err != nil {
		return fmt.Errorf("failed to create permissions: %w", err)
	}

	return nil
}

// getPermissionDescription returns description for each permission
func getPermissionDescription(permissionID int64) string {
	descriptions := map[int64]string{
		models.PermissionCreateUser:            "Create new users",
		models.PermissionReadUser:              "View user information",
		models.PermissionUpdateUser:            "Update user information",
		models.PermissionDeleteUser:            "Delete users",
		models.PermissionManageRoles:           "Manage user roles",
		models.PermissionViewOwnReport:         "View own reports",
		models.PermissionViewSubordinateReport: "View subordinate reports",
		models.PermissionViewAllReports:        "View all reports",
		models.PermissionCreateReport:          "Create new reports",
		models.PermissionDeleteReport:          "Delete reports",
		models.PermissionManageFinances:        "Manage financial operations",
		models.PermissionManageInventory:       "Manage inventory",
		models.PermissionPosition:              "Create positions",
		models.PermissionUpdatePosition:        "Update positions",
		models.PermissionDeletePosition:        "Delete positions",
		models.PermissionCreateDepartment:      "Create departments",
		models.PermissionUpdateDepartment:      "Update departments",
		models.PermissionDeleteDepartment:      "Delete departments",
		models.PermissionCreateCompany:         "Create companies",
		models.PermissionUpdateCompany:         "Update companies",
		models.PermissionDeleteCompany:         "Delete companies",
		models.PermissionCreateContract:        "Create contracts",
		models.PermissionUpdateContract:        "Update contracts",
		models.PermissionDeleteContract:        "Delete contracts",
		models.PermissionApproveRequests:       "Approve requests",
		models.PermissionCreateRequest:         "Create requests",
		models.PermissionUpdateRequest:         "Update requests",
		models.PermissionDeleteRequest:         "Delete requests",
	}
	return descriptions[permissionID]
}

func seedRolePermissions(db *gorm.DB) error {
	now := time.Now()

	// Build role_permissions array from RolePermissions map
	rolePermissions := make([]models.RolePermission, 0)
	for roleID, permissionIDs := range models.RolePermissions {
		for _, permissionID := range permissionIDs {
			rolePermissions = append(rolePermissions, models.RolePermission{
				RoleID:       roleID,
				PermissionID: permissionID,
				GrantedAt:    &now,
			})
		}
	}

	// Batch insert all role_permissions
	if err := db.Create(&rolePermissions).Error; err != nil {
		return fmt.Errorf("failed to create role_permissions: %w", err)
	}

	return nil
}

func seedUser(db *gorm.DB) error {
	pw, _ := utils.HashPassword("password123")

	users := []models.User{
		{
			FullName: "Super Admin",
			HashPassword: pw,
			Email: "super-admin@gmail.com",

		},
		{
			FullName: "Admin",
			HashPassword: pw,
			Email: "admin@gmail.com",

		},
		{
			FullName: "User",
			HashPassword: pw,
			Email: "user@gmail.com",

		},
	}

	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func seedUserRole(db *gorm.DB) error {
	now := time.Now()

	users := []models.UserRole{
		{
			UserID: 1,
			RoleID: 1,
			AssignedAt: &now,
		},
		{
			UserID: 2,
			RoleID: 2,
			AssignedAt: &now,

		},
		{
			UserID: 3,
			RoleID: 10,
			AssignedAt: &now,

		},
	}

	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("failed to create user role: %w", err)
	}

	return nil
}
