package services

import (
	"context"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/requests"
)

type CompanyRepo interface {
	CreateCompany(ctx context.Context, data *models.Company) error
	GetCompany(ctx context.Context, data map[string]interface{}) (*models.Company, error)
	GetAllCompaniesWithPagination(ctx context.Context, limit, offset int, data map[string]interface{}) ([]*models.Company, error)
	CountCompanies(ctx context.Context, data map[string]interface{}) (int64, error)
	UpdateCompany(ctx context.Context, id uint64, data map[string]interface{}) error
	DeleteCompany(ctx context.Context, id uint64) error
}

type companyService struct {
	repo CompanyRepo
}

func NewCompanyService(repo CompanyRepo) *companyService {
	return &companyService{repo: repo}
}

func (s *companyService) CreateCompany(ctx context.Context, data *requests.CreateCompanyRequest) (*models.Company, error) {
	company := &models.Company{
		SQLModel:    models.NewSQLModel(),
		Name:        data.Name,
		ParentID:    data.ParentID,
		Description: data.Description,
		FoundedDate: data.FoundedDate,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		Email:       data.Email,
	}

	if err := s.repo.CreateCompany(ctx, company); err != nil {
		return nil, common.ErrorCreateFailed.Clone().WrapErrorSafe(err)
	}

	return company, nil
}

func (s *companyService) FindByID(ctx context.Context, id uint64) (*models.Company, error) {
	company, err := s.repo.GetCompany(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *companyService) GetListCompaniesWithPagination(ctx context.Context, data requests.ListCompanyRequest) ([]*models.Company, error) {
	offset := (data.Page - 1) * data.Limit

	query := make(map[string]interface{})
	if data.ParentID != nil {
		query["parent_id"] = *data.ParentID
	}

	companies, err := s.repo.GetAllCompaniesWithPagination(ctx, data.Limit, offset, query)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (s *companyService) UpdateCompany(ctx context.Context, id uint64, data *requests.UpdateCompanyRequest) error {
	// Check if company exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("company not found")
	}

	// Build update map with only non-nil fields
	updates := make(map[string]interface{})
	if data.Name != nil {
		updates["name"] = *data.Name
	}
	if data.Description != nil {
		updates["description"] = *data.Description
	}
	if data.FoundedDate != nil {
		updates["founded_date"] = *data.FoundedDate
	}
	if data.Address != nil {
		updates["address"] = *data.Address
	}
	if data.PhoneNumber != nil {
		updates["phone_number"] = *data.PhoneNumber
	}
	if data.Email != nil {
		updates["email"] = *data.Email
	}

	if len(updates) == 0 {
		return common.ErrorValidation.Clone().WrapMessage("no fields to update")
	}

	if err := s.repo.UpdateCompany(ctx, id, updates); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}

func (s *companyService) DeleteCompany(ctx context.Context, id uint64) error {
	// Check if company exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("company not found")
	}

	if err := s.repo.DeleteCompany(ctx, id); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}
