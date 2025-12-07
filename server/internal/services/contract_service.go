package services

import (
	"context"
	"time"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/requests"
)

type ContractRepo interface {
	CreateContract(ctx context.Context, data *models.Contract) error
	GetContract(ctx context.Context, data map[string]interface{}) (*models.Contract, error)
	GetAllContractsWithPagination(ctx context.Context, limit, offset int, data map[string]interface{}) ([]*models.Contract, error)
	CountContracts(ctx context.Context, data map[string]interface{}) (int64, error)
	UpdateContract(ctx context.Context, id uint64, data map[string]interface{}) error
	DeleteContract(ctx context.Context, id uint64) error
}

type contractService struct {
	repo ContractRepo
}

func NewContractService(repo ContractRepo) *contractService {
	return &contractService{repo: repo}
}

func (s *contractService) CreateContract(ctx context.Context, data *requests.CreateContractRequest) (*models.Contract, error) {
	// Parse dates
	startDate, err := time.Parse("2006-01-02", data.StartDate)
	if err != nil {
		return nil, common.ErrorValidation.Clone().SetDetail("start_date", "invalid date format")
	}

	var endDate *time.Time
	if data.EndDate != nil {
		parsed, err := time.Parse("2006-01-02", *data.EndDate)
		if err != nil {
			return nil, common.ErrorValidation.Clone().SetDetail("end_date", "invalid date format")
		}
		endDate = &parsed
	}

	contract := &models.Contract{
		SQLModel:       models.NewSQLModel(),
		UserID:         data.UserID,
		CompanyID:      data.CompanyID,
		PositionID:     data.PositionID,
		ContractNumber: data.ContractNumber,
		ContractType:   models.ContractType(data.ContractType),
		StartDate:      startDate,
		EndDate:        endDate,
		Salary:         data.Salary,
		Status:         models.ContractStatus(data.Status),
		FilePath:       data.FilePath,
		Notes:          data.Notes,
	}

	if err := s.repo.CreateContract(ctx, contract); err != nil {
		return nil, common.ErrorCreateFailed.Clone().WrapErrorSafe(err)
	}

	return contract, nil
}

func (s *contractService) FindByID(ctx context.Context, id uint64) (*models.Contract, error) {
	contract, err := s.repo.GetContract(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return contract, nil
}

func (s *contractService) GetListContractsWithPagination(ctx context.Context, data requests.ListContractRequest) ([]*models.Contract, error) {
	offset := (data.Page - 1) * data.Limit

	query := make(map[string]interface{})
	if data.UserID != nil {
		query["user_id"] = *data.UserID
	}
	if data.CompanyID != nil {
		query["company_id"] = *data.CompanyID
	}
	if data.Status != nil {
		query["status"] = *data.Status
	}
	if data.Type != nil {
		query["contract_type"] = *data.Type
	}

	contracts, err := s.repo.GetAllContractsWithPagination(ctx, data.Limit, offset, query)
	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (s *contractService) UpdateContract(ctx context.Context, id uint64, data *requests.UpdateContractRequest) error {
	// Check if contract exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("contract not found")
	}

	// Build update map with only non-nil fields
	updates := make(map[string]interface{})
	if data.ContractType != nil {
		updates["contract_type"] = *data.ContractType
	}
	if data.StartDate != nil {
		parsed, err := time.Parse("2006-01-02", *data.StartDate)
		if err != nil {
			return common.ErrorValidation.Clone().SetDetail("start_date", "invalid date format")
		}
		updates["start_date"] = parsed
	}
	if data.EndDate != nil {
		updates["end_date"] = *data.EndDate
	}
	if data.Salary != nil {
		updates["salary"] = *data.Salary
	}
	if data.Status != nil {
		updates["status"] = *data.Status
	}
	if data.FilePath != nil {
		updates["file_path"] = *data.FilePath
	}
	if data.Notes != nil {
		updates["notes"] = *data.Notes
	}

	if len(updates) == 0 {
		return common.ErrorValidation.Clone().WrapMessage("no fields to update")
	}

	if err := s.repo.UpdateContract(ctx, id, updates); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}

func (s *contractService) DeleteContract(ctx context.Context, id uint64) error {
	// Check if contract exists
	_, err := s.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("contract not found")
	}

	if err := s.repo.DeleteContract(ctx, id); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}
