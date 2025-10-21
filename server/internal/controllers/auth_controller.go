package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/requests"
)

func LoginHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var lr requests.LoginRequest

		if err := c.BodyParser(&lr); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(common.ErrorBodyParser)
		}

		if err := lr.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(common.ErrorValidation.WrapDetail(err))
		}

		return c.Status(fiber.StatusOK).JSON(common.LoginSuccessful)
	}
}

func RegisterHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rr requests.LoginRequest

		if err := c.BodyParser(&rr); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(common.ErrorBodyParser)
		}

		if err := rr.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(common.ErrorValidation.WrapDetail(err))
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Check health successfuly",
		})
	}
}
