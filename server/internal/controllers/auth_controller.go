package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/repositories"
	"github.com/vlahanam/company-management/internal/requests"
	"github.com/vlahanam/company-management/internal/services"
)

func LoginHandler(db *gorm.DB, accessSecret, refreshSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var lr requests.LoginRequest

		if err := c.BodyParser(&lr); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := lr.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		es := services.NewUserService(rp)
		as := services.NewAuthService(es, accessSecret, refreshSecret)

		auth, err := as.Login(c.UserContext(), &lr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.LoginSuccessful.WrapData(auth))
	}
}

func RefreshHandler(db *gorm.DB, accessSecret, refreshSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rr requests.RefreshRequest

		if err := c.BodyParser(&rr); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rr.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		es := services.NewUserService(rp)
		as := services.NewAuthService(es, accessSecret, refreshSecret)

		auth, err := as.RefreshAccessToken(c.UserContext(), rr.RefreshToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.AccessTokenCreatedSuccessfully.WrapData(auth))
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
		sr := services.NewUserService(rp)

		if err := sr.CreateUser(c.UserContext(), &rr); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.RegisterSuccessful)
	}
}
