package authRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/controllers/authController"
	"github.com/subAlgo/userFiber/middleware"
)

func Handle(app *fiber.App, apiGroup fiber.Router) {
	//au := apiGroup.Group("/auth", middleware.EcoText)
	au := apiGroup.Group("/auth", middleware.EcoText)
	au.Post("/signup", authController.Signup)   // /api/auth/signup
	au.Post("/signin", authController.SignIn)   // /api/auth/signin
	au.Post("/signout", authController.SignOut) // /api/auth/signout
}
