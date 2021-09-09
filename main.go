package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/subAlgo/userFiber/database"
)

func main() {
	/*
		var err error
		if err = database.Connect(); err != nil {
			panic("could not connect to the database")
		}
		if err = database.Migration(); err != nil {
			log.Fatalln(err)
		}

	*/
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	//routes.Setup(app)
	//authRoute.Handler(app)
	app.Listen(":8000")
}
