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
	UpdateUser(ctx context.Context, id uint64, data map[string]interface{}) error
	DeleteUser(ctx context.Context, id uint64) error
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

func (es *userService) UpdateUser(ctx context.Context, id uint64, data *requests.UpdateUserRequest) error {
	// Check if user exists
	user, err := es.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("user not found")
	}

	// Check email uniqueness if email is being updated
	if data.Email != nil && *data.Email != user.Email {
		existingUser, _ := es.FindByEmail(ctx, *data.Email)
		if existingUser != nil {
			return common.ErrorValidation.Clone().SetDetail("email", "email already exists")
		}
	}

	// Build update map with only non-nil fields
	updates := make(map[string]interface{})
	if data.FullName != nil {
		updates["full_name"] = *data.FullName
	}
	if data.Email != nil {
		updates["email"] = *data.Email
	}
	if data.DateOfBirth != nil {
		updates["date_of_birth"] = *data.DateOfBirth
	}
	if data.Gender != nil {
		updates["gender"] = *data.Gender
	}
	if data.IdCardNumber != nil {
		updates["id_card_number"] = *data.IdCardNumber
	}
	if data.PhoneNumber != nil {
		updates["phone_number"] = *data.PhoneNumber
	}
	if data.Avatar != nil {
		updates["avatar"] = *data.Avatar
	}

	if len(updates) == 0 {
		return common.ErrorValidation.Clone().WrapMessage("no fields to update")
	}

	if err := es.er.UpdateUser(ctx, id, updates); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}

func (es *userService) DeleteUser(ctx context.Context, id uint64) error {
	// Check if user exists
	_, err := es.FindByID(ctx, id)
	if err != nil {
		return common.ErrorNotFound.Clone().WrapMessage("user not found")
	}

	if err := es.er.DeleteUser(ctx, id); err != nil {
		return common.ErrorInternal.Clone().WrapErrorSafe(err)
	}

	return nil
}
