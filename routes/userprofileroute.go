package routes

import (
	"strconv"

	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	data, dataErr := services.GetUserById(id)
	if dataErr != nil {
		return dataErr
	}

	return c.Render("userprofile", fiber.Map{
		"Title":       "User Profile",
		"Data":        data,
		"TrainerName": c.Locals("UserName"),
	}, "layout")
}
