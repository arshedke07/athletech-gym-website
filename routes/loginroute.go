package routes

import (
	"strings"
	"time"

	"github.com/arshedke07/athletech/services"
	"github.com/arshedke07/athletech/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func LoginUserRoute(c *fiber.Ctx, store *session.Store) error {
	if c.Method() == "GET" {
		return c.Render("loginUser", fiber.Map{
			"Title": "Login User",
		}, "layout")
	} else if c.Method() == "POST" {
		// sess, sessErr := store.Get(c)
		// if sessErr != nil {
		// 	return sessErr
		// }

		emailid := c.FormValue("email")
		password := c.FormValue("password")
		user, err := services.LoginUserService(emailid, password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid email or password",
			})
		}

		firstName := strings.ToUpper(string(user.FirstName[0])) + user.FirstName[1:]
		lastName := strings.ToUpper(string(user.LastName[0])) + user.LastName[1:]

		// sess.Set("UserId", user.UserId)
		// sess.Set("Name", firstName+" "+lastName)
		// sess.Save()

		// generate token after successful loginusing
		token, err := utils.GenerateToken(user.UserId, firstName+" "+lastName, "none")
		if err != nil {
			return err
		}

		// store the token in cookie
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour), // match token expiry
			HTTPOnly: true,                           // prevent JS access (protects from XSS)
			Secure:   true,                           // send only over HTTPS
			SameSite: "Strict",                       // or "Lax" to balance CSRF protection and UX
		})

		return c.Render("userhome", fiber.Map{
			"Title":    "User Home",
			"UserName": firstName + " " + lastName,
		}, "layout")
	}

	return nil
}

func LoginTrainerRoute(c *fiber.Ctx, store *session.Store) error {
	if c.Method() == "GET" {
		return c.Render("logintrainer", fiber.Map{
			"Title": "Login User",
		}, "layout")
	} else if c.Method() == "POST" {
		// sess, sessErr := store.Get(c)
		// if sessErr != nil {
		// 	return sessErr
		// }

		emailid := c.FormValue("email")
		password := c.FormValue("password")
		user, err := services.LoginTrainerService(emailid, password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid email or password",
			})
		}

		data, err := services.GetPendingUsers(user.Id)
		if err != nil {
			return err
		}

		firstName := strings.ToUpper(string(user.FirstName[0])) + user.FirstName[1:]
		lastName := strings.ToUpper(string(user.LastName[0])) + user.LastName[1:]

		token, err := utils.GenerateToken(user.Id, firstName+" "+lastName, "none")
		if err != nil {
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour), // match token expiry
			HTTPOnly: true,                           // prevent JS access (protects from XSS)
			Secure:   true,                           // send only over HTTPS
			SameSite: "Strict",                       // or "Lax" to balance CSRF protection and UX
		})

		// sess.Set("Name", firstName+" "+lastName)
		// sess.Set("TrainerId", user.Id)
		// sess.Save()

		return c.Render("trainerhome", fiber.Map{
			"Title":       "User Home",
			"Data":        data,
			"TrainerName": firstName + " " + lastName,
		}, "layout")
	}

	return nil
}
