package services

import (
	"context"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/requests"
)

type UserRepo interface {
	CreateUser(ctx context.Context, data *models.User) error
	GetFirst(ctx context.Context, data map[string]interface{}) (*models.User, error)
}

type userService struct {
	er UserRepo
}

func NewUserService(er UserRepo) *userService {
	return &userService{er: er}
}

func (es *userService) CreateUser(ctx context.Context, data *requests.RegisterRequest) error {
	emp, _ := es.FindByEmail(ctx, data.Email)
	if emp != nil {
		return common.ErrorValidation.Clone().SetDetail("email", models.ErrEmailAlreadyExists.Error())
	}

	hashPassword, err := common.HashPassword(data.Password)
	if err != nil {
		return common.ErrorCreateFailed.Clone().WrapError(err)
	}

	emp = &models.User{
		Email:        data.Email,
		FullName:     data.FullName,
		HashPassword: hashPassword,
	}

	if err := es.er.CreateUser(ctx, emp); err != nil {
		return common.ErrorCreateFailed.Clone().WrapErrorSafe(err)
	}

	return nil
}

func (es *userService) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	emp, err := es.er.GetFirst(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, err
	}

	return emp, nil
}
