package utils

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	userID string
	roles  []interface{}
}

var (
	ErrTokenMissingKey       = "AUTHORIZATION_TOKEN_MISSING"
	ErrInvalidTokenFormatKey = "INVALID_TOKEN_FORMAT"
	ErrTokenExpiredKey       = "INVALID_OR_EXPIRED_TOKEN"
	ErrPermissionDeniedKey   = "PERMISSION_DENIED"

	ErrTokenMissing       = errors.New("authorization token missing")
	ErrInvalidTokenFormat = errors.New("invalid token format")
	ErrTokenExpired       = errors.New("invalid or expired token")
	ErrPermissionDenied   = errors.New("you do not have permission to access this resource")
)

func AuthMiddleware(accessSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"key":   ErrTokenMissingKey,
				"error": ErrTokenMissing,
			})
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"key":   ErrInvalidTokenFormatKey,
				"error": ErrInvalidTokenFormat,
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(accessSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"key":   ErrTokenExpiredKey,
				"error": ErrTokenExpired,
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("userClaims", claims)
		}

		return c.Next()
	}
}

func CheckRole(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims := protectedHandler(c)

		userRoles := []string{}
		for _, v := range userClaims.roles {
			if str, ok := v.(string); ok {
				userRoles = append(userRoles, str)
			}
		}

		hasRole := false
		for _, ar := range allowedRoles {
			for _, ur := range userRoles {
				if ar == ur {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}

		if !hasRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"key":   ErrPermissionDeniedKey,
				"error": ErrPermissionDenied.Error(),
			})
		}

		return c.Next()
	}
}

func protectedHandler(c *fiber.Ctx) *UserClaims {
	claims := c.Locals("userClaims").(jwt.MapClaims)

	userID := claims["user_id"].(string)
	roles := claims["roles"].([]interface{})

	userClaims := &UserClaims{
		userID,
		roles,
	}

	return userClaims
}
