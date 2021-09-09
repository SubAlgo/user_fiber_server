package adminController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/response"
)

type changeUserDepartmentRequest struct {
	UserID       uint `json:"userID"`
	DepartmentID int  `json:"departmentID"`
}

func ChangeUserDepartment(c *fiber.Ctx) error {
	var user model.User
	var req changeUserDepartmentRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	database.DB.Find(&user, req.UserID)
	if user.ID == 0 {
		return response.Error(c, fiber.StatusNotFound, "Not found this user")
	}

	user.DepartmentID = req.DepartmentID
	if err := database.DB.Save(&user).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.JsonMessage(c, "Change user department success")
}
