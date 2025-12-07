package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/requests"
	"github.com/vlahanam/company-management/internal/repositories"
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
			c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.GetListSuccessResponse("users").WrapData(users))
	}
}

func GetUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Get user")
	}
}

func UpdateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Update user")
	}
}

func DeleteUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Delete user")
	}
}
