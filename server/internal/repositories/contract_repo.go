package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/vlahanam/company-management/internal/models"
)

func (s *mysqlStorage) CreateContract(ctx context.Context, data *models.Contract) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) GetContract(ctx context.Context, data map[string]interface{}) (*models.Contract, error) {
	var contract *models.Contract
	if err := s.db.WithContext(ctx).Where(data).First(&contract).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrContractNotFound
		}

		return nil, err
	}

	return contract, nil
}

func (s *mysqlStorage) GetAllContractsWithPagination(ctx context.Context, limit, offset int, data map[string]interface{}) ([]*models.Contract, error) {
	var contracts []*models.Contract

	qr := s.db.WithContext(ctx).Where(data)
	qr = qr.Limit(limit).Offset(offset)

	if err := qr.Find(&contracts).Error; err != nil {
		return nil, err
	}

	return contracts, nil
}

func (s *mysqlStorage) CountContracts(ctx context.Context, data map[string]interface{}) (int64, error) {
	var count int64

	if err := s.db.WithContext(ctx).Model(&models.Contract{}).Where(data).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *mysqlStorage) UpdateContract(ctx context.Context, id uint64, data map[string]interface{}) error {
	if err := s.db.WithContext(ctx).Model(&models.Contract{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) DeleteContract(ctx context.Context, id uint64) error {
	if err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Contract{}).Error; err != nil {
		return err
	}

	return nil
}
