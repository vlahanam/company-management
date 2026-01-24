package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/repositories"
	"github.com/vlahanam/company-management/internal/requests"
	"github.com/vlahanam/company-management/internal/services"
)

func GetListUsers(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rq requests.ListUserRequest

		if err := c.QueryParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorQueryParser)
		}

		rp := repositories.NewMySQLStorage(db)
		es := services.NewUserService(rp)

		users, err := es.GetListUsersWithPagination(c.UserContext(), rq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.GetListSuccessResponse("users").WrapData(users))
	}
}

func GetUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		es := services.NewUserService(rp)

		user, err := es.FindByID(c.UserContext(), uint64(uid.GetLocalID()))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorNotFound.Clone().WrapMessage("user not found"))
		}

		user.Mask(1) // User object type
		return c.Status(fiber.StatusOK).JSON(common.GetSuccessResponse("user").WrapData(user))
	}
}

func UpdateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		var rq requests.UpdateUserRequest
		if err := c.BodyParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rq.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		es := services.NewUserService(rp)

		if err := es.UpdateUser(c.UserContext(), uint64(uid.GetLocalID()), &rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.UpdateSuccessResponse("user"))
	}
}

func DeleteUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		es := services.NewUserService(rp)

		if err := es.DeleteUser(c.UserContext(), uint64(uid.GetLocalID())); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.DeleteSuccessResponse("user"))
	}
}
