package routers

import "github.com/gofiber/fiber/v2"

func PositionRoute(v fiber.Router) {
	v.Get("/positions", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "List position",
		})
	})
	v.Post("/positions", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Create position",
		})
	})
	v.Put("/positions/:position_id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Update position",
		})
	})
	v.Delete("/positions/:position_id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Delete position",
		})
	})
}