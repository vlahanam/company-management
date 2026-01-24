package services

import (
	"context"
	"time"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/requests"
)

type PermissionRepo interface {
	CreatePermission(ctx context.Context, data *models.Permission) error
	GetPermission(ctx context.Context, data map[string]interface{}) (*models.Permission, error)
	GetAllPermissionsWithPagination(ctx context.Context, limit, offset int) ([]*models.Permission, error)
	CountPermissions(ctx context.Context) (int64, error)
	UpdatePermission(ctx context.Context, id int64, data map[string]interface{}) error
	DeletePermission(ctx context.Context, id int64) error
}

type permissionService struct {
	repo PermissionRepo
}

func NewPermissionService(repo PermissionRepo) *permissionService {
	return &permissionService{repo: repo}
}

func (s *permissionService) CreatePermission(ctx context.Context, data *requests.CreatePermissionRequest) (*models.Permission, error) {
	now := time.Now().UTC()
	permission := &models.Permission{
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   &now,
	}

	if err := s.repo.CreatePermission(ctx, permission); err != nil {
		return nil, common.ErrorCreateFailed.Clone().WrapErrorSafe(err)
	}

	return permission, nil
}

func (s *permissionService) FindByID(ctx context.Context, id int64) (*models.Permission, error) {
	permission, err := s.repo.GetPermission(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *permissionService) GetListPermissionsWithPagination(ctx context.Context, data requests.ListPermissionRequest) ([]*models.Permission, error) {
	offset := (data.Page - 1) * data.Limit

	permissions, err := s.repo.GetAllPermissionsWithPagination(ctx, data.Limit, offset)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (s *permissionService) UpdatePermission(ctx context.Context, id int64, data *requests.UpdatePermissionRequest) error {
	// Check if permission exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("permission not found")
	}

	// Build update map with only non-nil fields
	updates := make(map[string]interface{})
	if data.Name != nil {
		updates["name"] = *data.Name
	}
	if data.Description != nil {
		updates["description"] = *data.Description
	}

	if len(updates) == 0 {
		return common.ErrorValidation.Clone().WrapMessage("no fields to update")
	}

	if err := s.repo.UpdatePermission(ctx, id, updates); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}

func (s *permissionService) DeletePermission(ctx context.Context, id int64) error {
	// Check if permission exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("permission not found")
	}

	if err := s.repo.DeletePermission(ctx, id); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}
