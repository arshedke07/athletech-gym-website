package routes

import (
	"strconv"

	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func GetUserProfile(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)

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
		"TrainerName": sess.Get("Name"),
	}, "layout")
}
