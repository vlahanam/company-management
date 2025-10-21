package routers

import "github.com/gofiber/fiber/v2"

func ContractRoute(v fiber.Router) {
	v.Get("/contracts", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "List contract",
		})
	})
	v.Post("/contracts", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Create contract",
		})
	})
	v.Put("/contracts/:contract_id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Update contract",
		})
	})
	v.Delete("/contracts/:contract_id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Delete contract",
		})
	})
}