package routes

import (
	"strconv"
	"time"

	"github.com/arshedke07/athletech/model"
	"github.com/arshedke07/athletech/services"
	"github.com/arshedke07/athletech/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func AddUser(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("signupform", fiber.Map{
			"Title": "Athletech",
		}, "layout")
	} else if c.Method() == "POST" {
		// hash the password and store it in database
		hashedPassword, err := utils.GeneratePassword(c.FormValue("password"))
		if err != nil {
			return err
		}

		user := model.AppUser{
			FirstName: c.FormValue("firstname"),
			LastName:  c.FormValue("lastname"),
			Password:  string(hashedPassword),
			Email:     c.FormValue("email"),
			Mobile:    c.FormValue("mobile"),
			Role:      "user",
		}

		ageStr := c.FormValue("age")
		age, _ := strconv.Atoi(ageStr)

		heightStr := c.FormValue("height")
		height, _ := strconv.Atoi(heightStr)

		daysStr := c.FormValue("days_available")
		days, _ := strconv.Atoi(daysStr)

		weightStr := c.FormValue("weight")
		weight, _ := strconv.Atoi(weightStr)

		profile := model.UserProfile{
			User:                  user,
			Age:                   age,
			Height:                height,
			Weight:                float32(weight),
			Gender:                c.FormValue("gender"),
			Experience:            c.FormValue("experience"),
			Goal:                  c.FormValue("goal"),
			CurrentBodyType:       c.FormValue("body_type"),
			GymAccess:             c.FormValue("gym_access"),
			DaysAvailable:         days,
			WorkoutTimePreference: c.FormValue("workout_time_preference"),
			DietaryRestrictions:   c.FormValue("dietary_restrictions"),
			Injuries:              c.FormValue("injuries"),
			MedicalConditions:     c.FormValue("medical_conditions"),
		}

		data, err := services.AddUserService(&profile)
		if err != nil {
			return err
		}
		// generate token after successful sign up
		token, tokenErr := utils.GenerateToken(data.UserId, data.FirstName+" "+data.LastName, data.Role)
		if tokenErr != nil {
			return tokenErr
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour), // match token expiry
			HTTPOnly: true,                           // prevent JS access (protects from XSS)
			Secure:   true,                           // send only over HTTPS
			SameSite: "Strict",                       // or "Lax" to balance CSRF protection and UX
		})

		// sess.Set("UserId", data.UserId)
		// sess.Set("Name", data.FirstName+" "+data.LastName)

		return c.Render("userhome", fiber.Map{
			"Title":    "Athletech",
			"UserName": data.FirstName + " " + data.LastName,
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

		hashedPassword, err := utils.GeneratePassword(c.FormValue("password"))
		if err != nil {
			return err
		}

		user := model.AppUser{
			FirstName: c.FormValue("firstname"),
			LastName:  c.FormValue("lastname"),
			Password:  string(hashedPassword),
			Email:     c.FormValue("email"),
			Mobile:    c.FormValue("mobile"),
			Role:      "trainer",
		}

		trainer := model.TrainerProfile{
			User:           user,
			Age:            age,
			Gender:         c.FormValue("gender"),
			Specialization: c.FormValue("specialization"),
			Experience:     c.FormValue("experience"),
			Languages:      c.FormValue("languages"),
			City:           c.FormValue("city"),
			State:          c.FormValue("state"),
			SocialMedia:    c.FormValue("social_media"),
			Description:    c.FormValue("description"),
		}

		data, err := services.AddTrainerService(&trainer)
		if err != nil {
			return err
		}

		token, tokenErr := utils.GenerateToken(data.UserId, data.FirstName+" "+data.LastName, data.Role)
		if tokenErr != nil {
			return tokenErr
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour), // match token expiry
			HTTPOnly: true,                           // prevent JS access (protects from XSS)
			Secure:   true,                           // send only over HTTPS
			SameSite: "Strict",                       // or "Lax" to balance CSRF protection and UX
		})

		return c.Render("trainerhome", fiber.Map{
			"Title":       "Athletech",
			"TrainerName": data.FirstName + " " + data.LastName,
		}, "layout")
	}
	return nil
}
