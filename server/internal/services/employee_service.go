package services

import (
	"context"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/requests"
)

type EmployeeRepo interface {
	CreateEmployee(ctx context.Context, data *models.Employee) error
	GetFirst(ctx context.Context, data map[string]interface{}) (*models.Employee, error)
}

type employeeService struct {
	er EmployeeRepo
}

func NewEmployeeService(er EmployeeRepo) *employeeService {
	return &employeeService{er: er}
}

func (es *employeeService) CreateEmployee(ctx context.Context, data *requests.RegisterRequest) error {
	emp, _ := es.FindByEmail(ctx, data.Email)
	if emp != nil {
		return models.ErrEmailAlreadyExists
	}

	hashPassword, err := common.HashPassword(data.Password)
	if err != nil {
		return err
	}

	emp = &models.Employee{
		Email:        data.Email,
		FullName:     data.FullName,
		HashPassword: hashPassword,
	}

	if err := es.er.CreateEmployee(ctx, emp); err != nil {
		return err
	}

	return nil
}

func (es *employeeService) FindByEmail(ctx context.Context, email string) (*models.Employee, error) {
	emp, err := es.er.GetFirst(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, err
	}

	return emp, nil
}
