package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/middleware"
	"github.com/subAlgo/userFiber/routes/adminRoute"
	"github.com/subAlgo/userFiber/routes/authRoute"
	"github.com/subAlgo/userFiber/routes/userRoute"
	"log"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})
	// apiGroup install main middleware
	apiGroup := app.Group("/api", middleware.CheckHeader, cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Content-Type",
	}))
	authRoute.Handle(app, apiGroup)
	userRoute.Handle(app, apiGroup)
	adminRoute.Handle(app, apiGroup)
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
