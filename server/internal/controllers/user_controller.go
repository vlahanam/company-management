package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetListUsers(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Get list user")
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
