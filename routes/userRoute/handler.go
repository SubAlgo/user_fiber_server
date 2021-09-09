package userRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/controllers/userController"
	"github.com/subAlgo/userFiber/middleware"
)

func Handle(app *fiber.App, apiGroup fiber.Router) {
	userGroup := apiGroup.Group("/user", middleware.CheckSignIn)
	userGroup.Get("/", userController.Profile)       // /api/user
	userGroup.Get("/:id", userController.GetUser)    // /api/user/{userID}
	userGroup.Get("/list", userController.GetUsers)  // /api/user/list
	userGroup.Post("/update", userController.Update) // /api/user/update
}
