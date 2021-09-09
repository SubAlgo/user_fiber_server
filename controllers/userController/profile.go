package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/response"
)

type getUserResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	RoleCode     string `json:"roleCode"`
	DepartmentID int    `json:"departmentID"`
}

func Profile(c *fiber.Ctx) error {
	var user model.User
	userID := c.Locals("userID")

	database.DB.Where("id = ?", userID).First(&user)
	res := getUserResponse{
		Name:         user.Name,
		Lastname:     user.Lastname,
		Email:        user.Email,
		RoleCode:     user.RoleCode,
		DepartmentID: user.DepartmentID,
	}
	return response.JSON(c, res)
}
