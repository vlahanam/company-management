package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/vlahanam/company-management/internal/models"
)

func (s *mysqlStorage) CreateUser(ctx context.Context, data *models.User) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) GetAllUserWithPagination(ctx context.Context, limit, offset int, data map[string]interface{}) ([]*models.User, error) {
	var emps []*models.User

	qr := s.db.WithContext(ctx).Where(data)
	qr = qr.Limit(limit).Offset(offset)

	if err := qr.Find(&emps).Error; err != nil {
		return nil, err
	}

	return emps, nil
}

func (s *mysqlStorage) CountDataByQuery(ctx context.Context, data map[string]interface{}) (int64, error) {
	var count int64

	if err := s.db.WithContext(ctx).Where(data).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *mysqlStorage) GetUser(ctx context.Context, data map[string]interface{}) (*models.User, error) {
	var emps *models.User
	if err := s.db.WithContext(ctx).Where(data).First(&emps).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrUserNotFound
		}

		return nil, err
	}

	return emps, nil
}

func (s *mysqlStorage) GetUserWithRole(ctx context.Context, data map[string]interface{}) (*models.User, error) {
	var emps *models.User
	if err := s.db.WithContext(ctx).Where(data).First(&emps).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrUserNotFound
		}

		return nil, err
	}

	return emps, nil
}

func (s *mysqlStorage) UpdateUser(ctx context.Context, id uint64, data map[string]interface{}) error {
	if err := s.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) DeleteUser(ctx context.Context, id uint64) error {
	if err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}
