package services

import (
	"context"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/requests"
)

type PositionRepo interface {
	CreatePosition(ctx context.Context, data *models.Position) error
	GetPosition(ctx context.Context, data map[string]interface{}) (*models.Position, error)
	GetPositionsByCompany(ctx context.Context, companyID uint64, limit, offset int) ([]*models.Position, error)
	CountPositions(ctx context.Context, data map[string]interface{}) (int64, error)
	UpdatePosition(ctx context.Context, id uint64, data map[string]interface{}) error
	DeletePosition(ctx context.Context, id uint64) error
}

type positionService struct {
	repo PositionRepo
}

func NewPositionService(repo PositionRepo) *positionService {
	return &positionService{repo: repo}
}

func (s *positionService) CreatePosition(ctx context.Context, companyID uint64, data *requests.CreatePositionRequest) (*models.Position, error) {
	position := &models.Position{
		SQLModel:    models.NewSQLModel(),
		CompanyID:   companyID,
		Name:        data.Name,
		Description: data.Description,
		Level:       data.Level,
	}

	if err := s.repo.CreatePosition(ctx, position); err != nil {
		return nil, common.ErrorCreateFailed.Clone().WrapErrorSafe(err)
	}

	return position, nil
}

func (s *positionService) FindByID(ctx context.Context, id uint64) (*models.Position, error) {
	position, err := s.repo.GetPosition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return position, nil
}

func (s *positionService) GetPositionsByCompanyWithPagination(ctx context.Context, companyID uint64, data requests.ListPositionRequest) ([]*models.Position, error) {
	offset := (data.Page - 1) * data.Limit

	positions, err := s.repo.GetPositionsByCompany(ctx, companyID, data.Limit, offset)
	if err != nil {
		return nil, err
	}

	return positions, nil
}

func (s *positionService) UpdatePosition(ctx context.Context, id uint64, data *requests.UpdatePositionRequest) error {
	// Check if position exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("position not found")
	}

	// Build update map with only non-nil fields
	updates := make(map[string]interface{})
	if data.Name != nil {
		updates["name"] = *data.Name
	}
	if data.Description != nil {
		updates["description"] = *data.Description
	}
	if data.Level != nil {
		updates["level"] = *data.Level
	}

	if len(updates) == 0 {
		return common.ErrorValidation.Clone().WrapMessage("no fields to update")
	}

	if err := s.repo.UpdatePosition(ctx, id, updates); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}

func (s *positionService) DeletePosition(ctx context.Context, id uint64) error {
	// Check if position exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("position not found")
	}

	if err := s.repo.DeletePosition(ctx, id); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}
