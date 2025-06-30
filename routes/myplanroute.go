package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func MyPlanRoute(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)
	return c.Render("myplan", fiber.Map{
		"Title":    "Athletech",
		"UserName": sess.Get("Name"),
	}, "layout")
}
