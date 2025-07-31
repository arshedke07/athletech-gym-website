package routes

import (
	"github.com/arshedke07/athletech/model"
	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func nilIfEmpty(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

func UserProgress(c *fiber.Ctx, store *session.Store) error {
	if c.Method() == "GET" {
		// sess, _ := store.Get(c)
		// id := sess.Get("UserId")
		id := c.Locals("UserId")
		val, ok := id.(int)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Could Not Convert UserId",
			})
		}

		data, start, err := services.GetUserProgress(val)
		if err != nil {
			return err
		}

		return c.Render("progresslog", fiber.Map{
			"Title":    "My Progress Log",
			"UserName": c.Locals("UserName"),
			"Data":     data,
			"Start":    start,
		}, "layout")
	} else if c.Method() == "POST" {
		// sess, _ := store.Get(c)
		// id := sess.Get("UserId")

		id := c.Locals("UserId")
		val, ok := id.(int)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Could Not Convert UserId",
			})
		}

		progress := model.Progress{
			CurrentWeight:  nilIfEmpty(c.FormValue("weight_current")),
			WeightGoal:     nilIfEmpty(c.FormValue("weight_goal")),
			CardioType:     nilIfEmpty(c.FormValue("cardio_type")),
			CurrentCardio:  nilIfEmpty(c.FormValue("current_cardio")),
			CardioGoal:     nilIfEmpty(c.FormValue("cardio_goal")),
			LiftType:       nilIfEmpty(c.FormValue("lift_type")),
			CurrentLift:    nilIfEmpty(c.FormValue("current_lift")),
			LiftGoal:       nilIfEmpty(c.FormValue("lift_goal")),
			CurrentBodyFat: nilIfEmpty(c.FormValue("current_body_fat")),
			BodyFatGoal:    nilIfEmpty(c.FormValue("body_fat_goal")),
		}

		userErr := services.UserProgressService(progress, val)
		if userErr != nil {
			return userErr
		}

		data, start, err := services.GetUserProgress(val)
		if err != nil {
			return err
		}

		return c.Render("progresslog", fiber.Map{
			"Title":    "User Progress",
			"Data":     data,
			"Start":    start,
			"UserName": c.Locals("UserName"),
		}, "layout")
	}

	return nil
}
