package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogoutRoute(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Render("home", fiber.Map{
		"Title": "Athletech",
	}, "layout")
}
