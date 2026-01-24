package repositories

import (
	"context"

	"github.com/vlahanam/company-management/internal/models"
)

func (s *mysqlStorage) GetUserRoleNames(ctx context.Context, userID uint64) ([]string, error) {
	var roleNames []string

	err := s.db.WithContext(ctx).
		Table("roles").
		Select("roles.name").
		Joins("INNER JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Pluck("name", &roleNames).Error

	if err != nil {
		return nil, err
	}

	return roleNames, nil
}

func (s *mysqlStorage) CreateRole(ctx context.Context, data *models.Role) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) GetRole(ctx context.Context, data map[string]interface{}) (*models.Role, error) {
	var role *models.Role
	if err := s.db.WithContext(ctx).Where(data).First(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (s *mysqlStorage) GetAllRolesWithPagination(ctx context.Context, limit, offset int) ([]*models.Role, error) {
	var roles []*models.Role

	qr := s.db.WithContext(ctx)
	qr = qr.Limit(limit).Offset(offset)

	if err := qr.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *mysqlStorage) CountRoles(ctx context.Context) (int64, error) {
	var count int64

	if err := s.db.WithContext(ctx).Model(&models.Role{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *mysqlStorage) UpdateRole(ctx context.Context, id int64, data map[string]interface{}) error {
	if err := s.db.WithContext(ctx).Model(&models.Role{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) DeleteRole(ctx context.Context, id int64) error {
	if err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Role{}).Error; err != nil {
		return err
	}

	return nil
}
