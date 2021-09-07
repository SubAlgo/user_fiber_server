package authRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/controllers/authController"
)

func Handle(app *fiber.App) {
	app.Post("api/signup", authController.Signup)
}
