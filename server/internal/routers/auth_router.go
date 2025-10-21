package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vlahanam/company-management/internal/controllers"
	"gorm.io/gorm"
)

func AuthRoute(v fiber.Router, db *gorm.DB) {
	v.Post("/login", controllers.LoginHandler())
	v.Post("/register", controllers.RegisterHandler(db))
}