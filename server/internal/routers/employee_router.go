package routers

import "github.com/gofiber/fiber/v2"

func EmployeeRoute(v fiber.Router) {
	v.Get("/employees", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "List employee",
		})
	})
	v.Post("/employees", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Create employee",
		})
	})
	v.Put("/employees/:employee_id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Update employee",
		})
	})
	v.Delete("/employees/:employee_id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Delete employee",
		})
	})
}