package routes

import (
	"github.com/gofiber/fiber/v2"
)

func MyPlanRoute(c *fiber.Ctx) error {
	userName := c.Locals("UserName")
	return c.Render("myplan", fiber.Map{
		"Title":    "Athletech",
		"UserName": userName,
	}, "layout")
}
