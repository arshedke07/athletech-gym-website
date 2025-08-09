package utils

import (
	"github.com/gofiber/fiber/v2"
)

func UserOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("Role")
		if role == nil || role.(string) != "user" {
			return c.Status(fiber.StatusUnauthorized).SendString("Access Denied: User only")
		}
		return c.Next()
	}
}

// TrainerOnly middleware
func TrainerOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("Role")
		if role == nil || role.(string) != "trainer" {
			return c.Status(fiber.StatusUnauthorized).SendString("Access Denied: Trainer only")
		}
		return c.Next()
	}
}
