package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/vlahanam/company-management/internal/models"
)

func (s *mysqlStorage) CreateCompany(ctx context.Context, data *models.Company) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) GetCompany(ctx context.Context, data map[string]interface{}) (*models.Company, error) {
	var company *models.Company
	if err := s.db.WithContext(ctx).Where(data).First(&company).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrCompanyNotFound
		}

		return nil, err
	}

	return company, nil
}

func (s *mysqlStorage) GetAllCompaniesWithPagination(ctx context.Context, limit, offset int, data map[string]interface{}) ([]*models.Company, error) {
	var companies []*models.Company

	qr := s.db.WithContext(ctx).Where(data)
	qr = qr.Limit(limit).Offset(offset)

	if err := qr.Find(&companies).Error; err != nil {
		return nil, err
	}

	return companies, nil
}

func (s *mysqlStorage) CountCompanies(ctx context.Context, data map[string]interface{}) (int64, error) {
	var count int64

	if err := s.db.WithContext(ctx).Model(&models.Company{}).Where(data).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *mysqlStorage) UpdateCompany(ctx context.Context, id uint64, data map[string]interface{}) error {
	if err := s.db.WithContext(ctx).Model(&models.Company{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) DeleteCompany(ctx context.Context, id uint64) error {
	if err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Company{}).Error; err != nil {
		return err
	}

	return nil
}
