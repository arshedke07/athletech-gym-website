package routes

import (
	"fmt"
	"strconv"

	"github.com/arshedke07/athletech/model"
	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func CreateDietRoute(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)

	if c.Method() == "GET" {
		id := c.Params("id")

		return c.Render("create_diet", fiber.Map{
			"Title":    "Athletech",
			"UserName": sess.Get("Name"),
			"UserId":   id,
		}, "layout")
	} else if c.Method() == "POST" {
		sess, _ := store.Get(c)
		id := c.Params("id")
		userId, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
		meals := []string{"breakfast", "morning_snack", "lunch", "pre_workout", "post_workout", "dinner"}

		for _, day := range days {
			protein, _ := strconv.Atoi(c.FormValue(day + "_protein"))
			fats, _ := strconv.Atoi(c.FormValue(day + "_fats"))
			carbs, _ := strconv.Atoi(c.FormValue(day + "_carbs"))
			calories, _ := strconv.Atoi(c.FormValue(day + "_calories"))

			nutrients := model.Nutrients{
				Protein:  protein,
				Fats:     fats,
				Carbs:    carbs,
				Calories: calories,
			}

			dietDay := model.DietDay{}
			dietDay.Day = day
			dietDay.Nutrients = nutrients

			for _, meal := range meals {
				field := fmt.Sprintf("%s_%s", day, meal)
				mealName := c.FormValue(field)

				if mealName == "" {
					continue
				}

				diet := model.Diet{
					Meal:     mealName,
					MealType: meal,
				}

				dietDay.Diet = append(dietDay.Diet, diet)
			}
			dietErr := services.CreateDietService(dietDay, userId)
			if dietErr != nil {
				return dietErr
			}
		}

		val := sess.Get("TrainerId") // interface{}
		str := fmt.Sprintf("%v", val)
		intVal, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Conversion failed:", err)
			// handle error
		}

		data, dataErr := services.GetPendingUsers(intVal)
		if dataErr != nil {
			return dataErr
		}

		return c.Render("trainerhome", fiber.Map{
			"Title":       "Athletech",
			"TrainerName": sess.Get("Name"),
			"Success":     true,
			"message":     "Diet Saved Successfully",
			"Data":        data,
		}, "layout")
	}
	return nil
}
