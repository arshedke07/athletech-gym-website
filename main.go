package main

import (
	"log"

	"github.com/arshedke07/athletech/routes"
	"github.com/arshedke07/athletech/services"
	"github.com/arshedke07/athletech/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views:                 engine,
		DisableStartupMessage: false,
	})

	// Initialize redis storage
	// redisStore := redis.New(redis.Config{
	// 	Host:     "localhost",
	// 	Port:     6379,
	// 	Password: "", // optional
	// 	Database: 0,  // optional
	// })

	// function to load	.env file into my go app
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// store := session.New(session.Config{
	// 	Storage: redisStore,
	// })
	// these are the basic routes accessible to all
	app.Get("/", routes.HomeRoute)
	app.Get("/signupform", routes.AddUser)
	app.Post("/signupform", routes.AddUser)

	app.Get("/loginUser", routes.LoginUserRoute)
	app.Post("/loginUser", routes.LoginUserRoute)
	app.Get("/logout", routes.LogoutRoute)

	app.Get("/trainerform", routes.AddTrainer)
	app.Post("/trainerform", routes.AddTrainer)

	// below these are all the user routes

	app.Get("/myplan", utils.Validate(), utils.UserOnly(), routes.MyPlanRoute)

	app.Get("/workoutplans", utils.Validate(), utils.UserOnly(), routes.WorkoutPlanRoute)

	app.Get("/dietplans", utils.Validate(), utils.UserOnly(), routes.DietPlanRoute)

	app.Get("/userhome", utils.Validate(), utils.UserOnly(), func(c *fiber.Ctx) error {
		UserName := c.Locals("UserName")

		return c.Render("userhome", fiber.Map{
			"UserName": UserName,
		}, "layout")
	})

	app.Get("/trainerselect", utils.Validate(), utils.UserOnly(), routes.TrainerSelectRoute)
	app.Post("/trainerselect", utils.Validate(), utils.UserOnly(), routes.TrainerSelectRoute)

	app.Get("/progresslog", utils.Validate(), utils.UserOnly(), routes.UserProgress)
	app.Post("/progresslog", utils.Validate(), utils.UserOnly(), routes.UserProgress)

	// below these are the trainer routes

	app.Get("/create_workout/:id", utils.Validate(), utils.TrainerOnly(), routes.CreateWorkoutRoute)
	app.Post("/create_workout/:id", utils.Validate(), utils.TrainerOnly(), routes.CreateWorkoutRoute)

	app.Get("/create_diet/:id", utils.Validate(), utils.TrainerOnly(), routes.CreateDietRoute)
	app.Post("/create_diet/:id", utils.Validate(), utils.TrainerOnly(), routes.CreateDietRoute)

	app.Get("/trainerhome", utils.Validate(), utils.TrainerOnly(), func(c *fiber.Ctx) error {
		id := c.Locals("UserId")
		val := id.(int)

		data, err := services.GetPendingUsers(val)
		if err != nil {
			return err
		}

		return c.Render("trainerhome", fiber.Map{
			"Title":       "Athletech",
			"Data":        data,
			"TrainerName": c.Locals("UserName"),
		}, "layout")
	})

	app.Get("/userprofile/:id", utils.Validate(), utils.TrainerOnly(), routes.GetUserProfile)

	app.Listen(":3000")
}
