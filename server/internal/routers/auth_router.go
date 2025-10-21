package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vlahanam/company-management/internal/controllers"
)

func AuthRoute(v fiber.Router) {
	v.Post("/login", controllers.LoginHandler())
	v.Post("/register", controllers.RegisterHandler())
}