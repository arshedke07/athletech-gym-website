package routes

import (
	"errors"
	"fmt"

	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
)

func DietPlanRoute(c *fiber.Ctx) error {
	id := c.Locals("UserId")
	val, ok := id.(int)
	if !ok {
		return errors.New("variable is not of type int")
	}

	data, err := services.GetDietService(val)
	if data == nil || len(*data) == 0 {
		fmt.Println("No data")
		return c.Render("diet_plans", fiber.Map{
			"Title":   "Athletech",
			"Message": "No Meal Plan to Display",
		}, "layout")
	}
	if err != nil {
		return err
	}

	return c.Render("diet_plans", fiber.Map{
		"Title":       "Athletech",
		"Data":        data,
		"TrainerName": c.Locals("UserId"),
	}, "layout")
}
