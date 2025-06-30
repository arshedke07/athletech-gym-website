package main

import (
	"github.com/arshedke07/athletech/routes"
	"github.com/arshedke07/athletech/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views:                 engine,
		DisableStartupMessage: false,
	})

	// Initialize redis storage
	redisStore := redis.New(redis.Config{
		Host:     "localhost",
		Port:     6379,
		Password: "", // optional
		Database: 0,  // optional
	})

	store := session.New(session.Config{
		Storage: redisStore,
	})

	app.Get("/", routes.HomeRoute)
	app.Get("/signupform", func(c *fiber.Ctx) error {
		return routes.AddUser(c, *store)
	})
	app.Post("/signupform", func(c *fiber.Ctx) error {
		return routes.AddUser(c, *store)
	})
	app.Get("/trainerform", routes.AddTrainer)
	app.Post("/trainerform", routes.AddTrainer)
	app.Get("/myplan", func(c *fiber.Ctx) error {
		return routes.MyPlanRoute(c, store)
	})
	app.Get("/workoutplans", func(c *fiber.Ctx) error {
		return routes.WorkoutPlanRoute(c, store)
	})
	app.Get("/dietplans", func(c *fiber.Ctx) error {
		return routes.DietPlanRoute(c, store)
	})
	app.Get("/loginUser", func(c *fiber.Ctx) error {
		return routes.LoginUserRoute(c, store)
	})
	app.Post("/loginUser", func(c *fiber.Ctx) error {
		return routes.LoginUserRoute(c, store)
	})
	app.Get("/loginTrainer", func(c *fiber.Ctx) error {
		return routes.LoginTrainerRoute(c, store)
	})
	app.Post("/loginTrainer", func(c *fiber.Ctx) error {
		return routes.LoginTrainerRoute(c, store)
	})
	app.Get("/create_workout/:id", func(c *fiber.Ctx) error {
		return routes.CreateWorkoutRoute(c, store)
	})
	app.Post("/create_workout/:id", func(c *fiber.Ctx) error {
		return routes.CreateWorkoutRoute(c, store)
	})
	app.Get("/create_diet/:id", func(c *fiber.Ctx) error {
		return routes.CreateDietRoute(c, store)
	})
	app.Post("/create_diet/:id", func(c *fiber.Ctx) error {
		return routes.CreateDietRoute(c, store)
	})
	app.Get("/trainerselect", func(c *fiber.Ctx) error {
		return routes.TrainerSelectRoute(c, store)
	})
	app.Post("/trainerselect", func(c *fiber.Ctx) error {
		return routes.TrainerSelectRoute(c, store)
	})
	app.Get("/userhome", func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		name := sess.Get("Name")

		successMsg := sess.Get("success")
		errorMsg := sess.Get("error")

		// Clear the messages after reading
		sess.Delete("success")
		sess.Delete("error")
		sess.Save()

		return c.Render("userhome", fiber.Map{
			"Success":  successMsg,
			"Error":    errorMsg,
			"UserName": name,
		}, "layout")
	})
	app.Get("/trainerhome", func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)

		id := sess.Get("TrainerId")
		val := id.(int)

		data, err := services.GetPendingUsers(val)
		if err != nil {
			return err
		}

		return c.Render("trainerhome", fiber.Map{
			"Title":       "Athletech",
			"Data":        data,
			"TrainerName": sess.Get("Name"),
		}, "layout")
	})

	app.Get("/userprofile/:id", func(c *fiber.Ctx) error {
		return routes.GetUserProfile(c, store)
	})

	app.Listen(":3000")
}
