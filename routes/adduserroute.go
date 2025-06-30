package routes

import (
	"strconv"

	"github.com/arshedke07/athletech/model"
	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
)

func AddUser(c *fiber.Ctx, store session.Store) error {
	if c.Method() == "GET" {
		return c.Render("signupform", fiber.Map{
			"Title": "Athletech",
		}, "layout")
	} else if c.Method() == "POST" {
		sess, _ := store.Get(c)

		user := model.AppUser{
			FirstName: c.FormValue("firstname"),
			LastName:  c.FormValue("lastname"),
			Password:  c.FormValue("password"),
			Email:     c.FormValue("email"),
			Mobile:    c.FormValue("mobile"),
		}

		profile := model.Preferences{
			Age:                   c.FormValue("age"),
			Height:                c.FormValue("height"),
			Weight:                c.FormValue("weight"),
			Gender:                c.FormValue("gender"),
			Experience:            c.FormValue("experience"),
			Goal:                  c.FormValue("goal"),
			CurrentBodyType:       c.FormValue("body_type"),
			GymAccess:             c.FormValue("gym_access"),
			DaysAvailable:         c.FormValue("days_available"),
			WorkoutTimePreference: c.FormValue("workout_time_preference"),
			DietaryRestrictions:   c.FormValue("dietary_restrictions"),
			Injuries:              c.FormValue("injuries"),
			MedicalConditions:     c.FormValue("medical_conditions"),
		}

		data, err := services.AddUserService(&user, &profile)
		if err != nil {
			return err
		}

		sess.Set("UserId", data.UserId)
		sess.Set("Name", data.FirstName+" "+data.LastName)

		return c.Render("userhome", fiber.Map{
			"Title":    "Athletech",
			"UserName": sess.Get("Name"),
		}, "layout")
	}
	return nil
}

func AddTrainer(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("trainerform", fiber.Map{
			"Title": "Athletech",
		}, "layout")
	} else if c.Method() == "POST" {
		age, _ := strconv.Atoi(c.FormValue("age"))

		user := model.Trainer{
			FirstName:      c.FormValue("firstname"),
			LastName:       c.FormValue("lastname"),
			Age:            age,
			Qualifications: c.FormValue("qualifications"),
			Gender:         c.FormValue("gender"),
			Password:       c.FormValue("password"),
			Email:          c.FormValue("email"),
			Mobile:         c.FormValue("mobile"),
		}

		data, err := services.AddTrainerService(&user)
		if err != nil {
			return err
		}

		return c.Render("trainerhome", fiber.Map{
			"Title":       "Athletech",
			"TrainerName": data.FirstName + " " + data.LastName,
		}, "layout")
	}
	return nil
}
