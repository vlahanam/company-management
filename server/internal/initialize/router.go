package initialize

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vlahanam/company-management/internal/routers"
)

func InitRoute(cfg *Config) {
	app := fiber.New()

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Check health successfuly",
		})
	})

	v1 := app.Group("api/v1")

	routers.AuthRoute(v1)
	routers.CompanyRoute(v1)
	routers.ContractRoute(v1)
	routers.EmployeeRoute(v1)
	routers.PositionRoute(v1)

	port := ":" + cfg.Fiber.Port
	app.Listen(port)
}
