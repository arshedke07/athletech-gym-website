package routes

import (
	"strings"
	"time"

	"github.com/arshedke07/athletech/services"
	"github.com/arshedke07/athletech/utils"
	"github.com/gofiber/fiber/v2"
)

func LoginUserRoute(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("loginUser", fiber.Map{
			"Title": "Login User",
		}, "layout")
	} else if c.Method() == "POST" {
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

		// generate token after successful loginusing
		token, err := utils.GenerateToken(user.UserId, firstName+" "+lastName, user.Role)
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

		if user.Role == "user" {
			return c.Render("userhome", fiber.Map{
				"Title":    "User Home",
				"UserName": firstName + " " + lastName,
			}, "layout")
		} else if user.Role == "trainer" {
			users, err := services.GetPendingUsers(user.UserId)
			if err != nil {
				return err
			}
			return c.Render("trainerhome", fiber.Map{
				"Title":       "User Home",
				"TrainerName": firstName + " " + lastName,
				"Data":        users,
			}, "layout")
		}
	}

	return nil
}
