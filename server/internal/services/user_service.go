package services

import (
	"context"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/requests"
	"github.com/vlahanam/company-management/utils"
)

type UserRepo interface {
	CreateUser(ctx context.Context, data *models.User) error
	GetUser(ctx context.Context, data map[string]interface{}) (*models.User, error)
	GetUserRoleNames(ctx context.Context, userID uint64) ([]string, error)
	CountDataByQuery(ctx context.Context, data map[string]interface{}) (int64, error)
	GetAllUserWithPagination(ctx context.Context, limit, offset int, data map[string]interface{}) ([]*models.User, error)
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

	hashPassword, err := utils.HashPassword(data.Password)
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
	emp, err := es.er.GetUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (es *userService) FindByID(ctx context.Context, id uint64) (*models.User, error) {
	emp, err := es.er.GetUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (es *userService) GetUser(ctx context.Context, data map[string]interface{}) (*models.User, error) {
	emp, err := es.er.GetUser(ctx, data)
	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (es *userService) GetRoleNamesByUserID(ctx context.Context, userID uint64) ([]string, error) {
	return es.er.GetUserRoleNames(ctx, userID)
}

func (es *userService) GetListUsersWithPagination(ctx context.Context, data requests.ListUserRequest) ([]*models.User, error) {
	offset := (data.Page - 1) * data.Limit

	emp, err := es.er.GetAllUserWithPagination(ctx, data.Limit, offset, map[string]interface{}{
		"123": 123,
	})
	
	if err != nil {
		return nil, err
	}

	return emp, nil
}
