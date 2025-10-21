package routers

import "github.com/gofiber/fiber/v2"

func CompanyRoute(v fiber.Router) {
	v.Get("/companies", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "List company",
		})
	})
	v.Post("/companies", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Create company",
		})
	})
	v.Put("/companies/:company_id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Update company",
		})
	})
	v.Delete("/companies/:company_id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Delete company",
		})
	})
}