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

func (s *mysqlStorage) GetAllUser(ctx context.Context, data map[string]interface{}) ([]models.User, error) {
	var emps []models.User
	if err := s.db.WithContext(ctx).Where(data).Find(&emps).Error; err != nil {
		return nil, err
	}

	return emps, nil
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
