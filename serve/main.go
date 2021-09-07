package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/subAlgo/userFiber/controllers/authController"
	"github.com/subAlgo/userFiber/database"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	//routes.Setup(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})
	app.Post("/api/signup", authController.Signup)
	app.Listen(":3000")
}

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}
