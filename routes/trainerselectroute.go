package routes

import (
	"fmt"
	"strconv"

	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Payload struct {
	Id string `json:"trainer_id"`
}

func TrainerSelectRoute(c *fiber.Ctx, store *session.Store) error {
	if c.Method() == "GET" {
		sess, _ := store.Get(c)

		data, err := services.GetAllTrainers()
		if err != nil {
			return err
		}
		return c.Render("trainerselect", fiber.Map{
			"Title":    "Select Trainer of your Choice",
			"Data":     data,
			"UserName": sess.Get("Name"),
		}, "layout")
	} else if c.Method() == "POST" {
		sess, err := store.Get(c)
		fmt.Println(err)

		id := sess.Get("UserId")
		fmt.Println(id)
		str := fmt.Sprintf("%v", id) // converted id variable to string and then convert to int
		i, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Conversion error:", err)
			return err
		}

		var data Payload

		parseErr := c.BodyParser(&data)
		if parseErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		trainerId, _ := strconv.Atoi(data.Id)
		fmt.Println(trainerId)

		userErr := services.UpdateUserService(i, trainerId)
		if userErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Success": false,
				"Message": "An error occurred while selecting trainer.",
			})
		}

		return c.JSON(fiber.Map{
			"Title":    "Athletech",
			"Success":  true,
			"Message":  "Trainer Selected Successfully",
			"UserName": sess.Get("Name"),
		}, "layout")
	}
	return nil
}
