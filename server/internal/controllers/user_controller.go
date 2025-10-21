package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetListUsers(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func GetUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func UpdateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func DeleteUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
