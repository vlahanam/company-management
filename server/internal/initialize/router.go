package initialize

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoute(cfg *Config) {
	app := fiber.New()

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Check health successfuly",
		})
	})

	port := ":" + cfg.Fiber.Port
	app.Listen(port)
}
