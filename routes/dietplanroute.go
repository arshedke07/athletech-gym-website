package routes

import (
	"errors"

	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DietPlanRoute(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)
	id := sess.Get("UserId")
	val, ok := id.(int)
	if !ok {
		return errors.New("variable is not of type int")
	}

	data, err := services.GetDietService(val)
	if err != nil {
		return err
	}

	return c.Render("diet_plans", fiber.Map{
		"Title":       "Athletech",
		"Data":        data,
		"TrainerName": sess.Get("Name"),
	}, "layout")
}
