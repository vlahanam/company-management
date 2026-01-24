package initialize

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"

	"github.com/vlahanam/company-management/internal/controllers"
	"github.com/vlahanam/company-management/internal/models"
	"github.com/vlahanam/company-management/utils"
)

func InitRoute(cfg *Config, db *gorm.DB) {
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.AllowedOrigins,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Check health successfuly",
		})
	})

	v1 := app.Group("api/v1")

	v1.Post("/login", controllers.LoginHandler(db, cfg.Auth.AccessSecret, cfg.Auth.RefreshSecret))
	v1.Post("/refresh", controllers.RefreshHandler(db, cfg.Auth.AccessSecret, cfg.Auth.RefreshSecret))
	v1.Post("/register", controllers.RegisterHandler(db))

	v1.Use(utils.AuthMiddleware(cfg.Auth.AccessSecret))

	v1.Get("/users", utils.CheckRole([]string{models.RoleNames[models.RoleSuperAdmin]}), controllers.GetListUsers(db))
	v1.Get("/users/:id", controllers.GetUser(db))
	v1.Put("/users/:id", controllers.UpdateUser(db))
	v1.Delete("/users/:id", controllers.DeleteUser(db))

	v1.Post("/companies", controllers.CreateCompany(db))
	v1.Get("/companies", controllers.GetListCompanies(db))
	v1.Get("/companies/:id", controllers.GetCompany(db))
	v1.Put("/companies/:id", controllers.UpdateCompany(db))
	v1.Delete("/companies/:id", controllers.DeleteCompany(db))

	v1.Post("/positions/:company_id", controllers.CreatePosition(db))
	v1.Get("/positions/:company_id", controllers.GetListPositions(db))
	v1.Get("/positions/:id", controllers.GetPosition(db))
	v1.Put("/positions/:id", controllers.UpdatePosition(db))
	v1.Delete("/positions/:id", controllers.DeletePosition(db))

	v1.Post("/contracts", controllers.CreateContract(db))
	v1.Get("/contracts", controllers.GetListContracts(db))
	v1.Get("/contracts/:id", controllers.GetContract(db))
	v1.Put("/contracts/:id", controllers.UpdateContract(db))
	v1.Delete("/contracts/:id", controllers.DeleteContract(db))

	v1.Post("/roles", controllers.CreateRole(db))
	v1.Get("/roles", controllers.GetListRoles(db))
	v1.Get("/roles/:id", controllers.GetRole(db))
	v1.Put("/roles/:id", controllers.UpdateRole(db))
	v1.Delete("/roles/:id", controllers.DeleteRole(db))

	v1.Post("/permissions", controllers.CreatePermission(db))
	v1.Get("/permissions", controllers.GetListPermissions(db))
	v1.Get("/permissions/:id", controllers.GetPermission(db))
	v1.Put("/permissions/:id", controllers.UpdatePermission(db))
	v1.Delete("/permissions/:id", controllers.DeletePermission(db))

	port := ":" + cfg.Fiber.Port
	app.Listen(port)
}
