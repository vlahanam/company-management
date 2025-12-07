package services

import (
	"context"
	"time"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/requests"
)

type RoleRepo interface {
	CreateRole(ctx context.Context, data *models.Role) error
	GetRole(ctx context.Context, data map[string]interface{}) (*models.Role, error)
	GetAllRolesWithPagination(ctx context.Context, limit, offset int) ([]*models.Role, error)
	CountRoles(ctx context.Context) (int64, error)
	UpdateRole(ctx context.Context, id int64, data map[string]interface{}) error
	DeleteRole(ctx context.Context, id int64) error
}

type roleService struct {
	repo RoleRepo
}

func NewRoleService(repo RoleRepo) *roleService {
	return &roleService{repo: repo}
}

func (s *roleService) CreateRole(ctx context.Context, data *requests.CreateRoleRequest) (*models.Role, error) {
	now := time.Now().UTC()
	role := &models.Role{
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   &now,
	}

	if err := s.repo.CreateRole(ctx, role); err != nil {
		return nil, common.ErrorCreateFailed.Clone().WrapErrorSafe(err)
	}

	return role, nil
}

func (s *roleService) FindByID(ctx context.Context, id int64) (*models.Role, error) {
	role, err := s.repo.GetRole(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *roleService) GetListRolesWithPagination(ctx context.Context, data requests.ListRoleRequest) ([]*models.Role, error) {
	offset := (data.Page - 1) * data.Limit

	roles, err := s.repo.GetAllRolesWithPagination(ctx, data.Limit, offset)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *roleService) UpdateRole(ctx context.Context, id int64, data *requests.UpdateRoleRequest) error {
	// Check if role exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("role not found")
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

	if err := s.repo.UpdateRole(ctx, id, updates); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}

func (s *roleService) DeleteRole(ctx context.Context, id int64) error {
	// Check if role exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("role not found")
	}

	if err := s.repo.DeleteRole(ctx, id); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}
