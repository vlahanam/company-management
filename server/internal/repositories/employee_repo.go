package repositories

import (
	"context"
	"errors"

	"github.com/vlahanam/company-management/internal/models"
	"gorm.io/gorm"
)

func (s *mysqlStorage) CreateEmployee(ctx context.Context, data *models.Employee) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) GetAll(ctx context.Context, data map[string]interface{}) ([]models.Employee, error) {
	var emps []models.Employee
	if err := s.db.WithContext(ctx).Where(data).Find(&emps).Error; err != nil {
		return nil, err
	}

	return emps, nil
}

func (s *mysqlStorage) GetFirst(ctx context.Context, data map[string]interface{}) (*models.Employee, error) {
	var emps *models.Employee
	if err := s.db.WithContext(ctx).Where(data).First(&emps).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrEmployeeNotFound
		}

		return nil, err
	}

	return emps, nil
}