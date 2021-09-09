package adminRoute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/controllers/adminController"
	"github.com/subAlgo/userFiber/middleware"
)

func Handle(app *fiber.App, apiGroup fiber.Router) {
	adminGroup := apiGroup.Group("/admin", middleware.CheckSignIn, middleware.CheckIsAdmin)
	adminGroup.Post("/department", adminController.CreateDepartment)                 // /api/admin/department
	adminGroup.Post("/update_user_department", adminController.ChangeUserDepartment) // /api/admin/update_user_department
}
