package utils

import (
	"fmt"
	"os"

	"github.com/arshedke07/athletech/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// function to validate user before letting them access any protected route on the website
func Validate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies("jwt")
		// parse the token from the cookie and store the payload in custom claims created
		// the keyfunc function returns the secret key signature to validate the token
		token, err := jwt.ParseWithClaims(tokenStr, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			// fmt.Println("Algorithm:", token.Header["alg"])

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// fmt.Println("Env variable", os.Getenv("JWT_SECRET"))
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or Expired Token Please Try Logging In Again",
			})
		}
		// extract the claims payload from the token
		claims, ok := token.Claims.(*model.Claims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Could not parse token claims",
			})
		}

		// store the required data from the token payload and store in c.Locals
		// it allows middleware and route handlers share data from the same request
		// Lives only during one HTTP request (deleted after response is sent) and cannot access it across different requests/sessions
		c.Locals("UserId", claims.UserId)
		c.Locals("UserName", claims.UserName)

		return c.Next()
	}
}
