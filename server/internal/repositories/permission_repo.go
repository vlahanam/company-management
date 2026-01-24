package repositories

import (
	"context"

	"github.com/vlahanam/company-management/internal/models"
)

func (s *mysqlStorage) CreatePermission(ctx context.Context, data *models.Permission) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) GetPermission(ctx context.Context, data map[string]interface{}) (*models.Permission, error) {
	var permission *models.Permission
	if err := s.db.WithContext(ctx).Where(data).First(&permission).Error; err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *mysqlStorage) GetAllPermissionsWithPagination(ctx context.Context, limit, offset int) ([]*models.Permission, error) {
	var permissions []*models.Permission

	qr := s.db.WithContext(ctx)
	qr = qr.Limit(limit).Offset(offset)

	if err := qr.Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

func (s *mysqlStorage) CountPermissions(ctx context.Context) (int64, error) {
	var count int64

	if err := s.db.WithContext(ctx).Model(&models.Permission{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *mysqlStorage) UpdatePermission(ctx context.Context, id int64, data map[string]interface{}) error {
	if err := s.db.WithContext(ctx).Model(&models.Permission{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) DeletePermission(ctx context.Context, id int64) error {
	if err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Permission{}).Error; err != nil {
		return err
	}

	return nil
}
