package routes

import (
	"fmt"

	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func WorkoutPlanRoute(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)
	id := sess.Get("UserId")
	value, ok := id.(int)
	if !ok {
		fmt.Println("variable is not an int")
	}

	data, err := services.GetWokoutService(value)
	if err != nil {
		return err
	}

	return c.Render("workout_plans", fiber.Map{
		"Title":       "Workout Plans",
		"Data":        data,
		"TrainerName": sess.Get("Name"),
	}, "layout")
}
