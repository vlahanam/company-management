package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get token from Authorization header
	authHeader := c.Get("Authorization")

	// Check if token is present
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token missing",
		})
	}

	// Check if the token format is correct (Bearer <token>)
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token format",
		})
	}

	// Extract token (strip "Bearer " prefix)
	token := authHeader[len("Bearer "):]

	// Validate the token (you can implement your own logic here)
	if token != "valid-token" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Continue to the next handler if authentication is successful
	return c.Next()
}
