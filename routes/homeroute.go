package routes

import "github.com/gofiber/fiber/v2"

func HomeRoute(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Athletech",
	}, "layout")
}
