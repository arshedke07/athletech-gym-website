package routes

import (
	"fmt"
	"strconv"

	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func CreateWorkoutRoute(c *fiber.Ctx, store *session.Store) error {
	if c.Method() == "GET" {
		id := c.Params("id")

		return c.Render("create_workout", fiber.Map{
			"Title": "Create Workout Plan For UserName",
			"Id":    id,
		}, "layout")
	} else if c.Method() == "POST" {
		sess, _ := store.Get(c)
		idStr := c.Params("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"Success": false,
				"message": "invalid user id",
			})
		}

		var exercises [][]string
		var desc [][]string
		var muscles []string
		days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
			return c.Status(400).JSON(fiber.Map{
				"Success": false,
				"message": "Failed to parse form",
			})
		}

		for i := 0; i < 7; i++ {
			day := days[i]
			arg := day + "_exercise[]"
			arg2 := day + "_muscles"

			field := form.Value[arg]
			muscle := ""
			if val, ok := form.Value[arg2]; ok && len(val) > 0 {
				muscle = val[0]
			}

			exercises = append(exercises, field)
			muscles = append(muscles, muscle)
		}

		for i := 0; i < 7; i++ {
			day := days[i]
			arg := day + "_description[]"
			field := form.Value[arg]

			desc = append(desc, field)
		}

		newErr := services.CreateWorkoutService(id, exercises, desc, muscles)
		if newErr != nil {
			fmt.Println("Error occured", newErr.Error())
			return c.JSON(fiber.Map{
				"Success": false,
				"message": "Error creating workout",
			})
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
			"message":     "Workout Saved Successfully",
			"Data":        data,
		}, "layout")
	}
	return nil
}
