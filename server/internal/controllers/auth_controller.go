package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/internal/repositories"
	"github.com/vlahanam/company-management/internal/requests"
	"github.com/vlahanam/company-management/internal/services"
	"gorm.io/gorm"
)

func LoginHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var lr requests.LoginRequest

		if err := c.BodyParser(&lr); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := lr.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		return c.Status(fiber.StatusOK).JSON(common.LoginSuccessful)
	}
}

func RegisterHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rr requests.RegisterRequest

		if err := c.BodyParser(&rr); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rr.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		sr := services.NewEmployeeService(rp)

		if err := sr.CreateEmployee(c.UserContext(), &rr); err != nil {
			if errors.Is(err, models.ErrEmailAlreadyExists) {
				return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.SetDetail("email", err.Error()))
			}

			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorCreateFailed)
		}

		return c.Status(fiber.StatusOK).JSON(common.RegisterSuccessful)
	}
}
