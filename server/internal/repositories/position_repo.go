package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/vlahanam/company-management/internal/models"
)

func (s *mysqlStorage) CreatePosition(ctx context.Context, data *models.Position) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) GetPosition(ctx context.Context, data map[string]interface{}) (*models.Position, error) {
	var position *models.Position
	if err := s.db.WithContext(ctx).Where(data).First(&position).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrPositionNotFound
		}

		return nil, err
	}

	return position, nil
}

func (s *mysqlStorage) GetPositionsByCompany(ctx context.Context, companyID uint64, limit, offset int) ([]*models.Position, error) {
	var positions []*models.Position

	qr := s.db.WithContext(ctx).Where("company_id = ?", companyID)
	qr = qr.Limit(limit).Offset(offset)

	if err := qr.Find(&positions).Error; err != nil {
		return nil, err
	}

	return positions, nil
}

func (s *mysqlStorage) CountPositions(ctx context.Context, data map[string]interface{}) (int64, error) {
	var count int64

	if err := s.db.WithContext(ctx).Model(&models.Position{}).Where(data).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *mysqlStorage) UpdatePosition(ctx context.Context, id uint64, data map[string]interface{}) error {
	if err := s.db.WithContext(ctx).Model(&models.Position{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *mysqlStorage) DeletePosition(ctx context.Context, id uint64) error {
	if err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Position{}).Error; err != nil {
		return err
	}

	return nil
}
