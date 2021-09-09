package userController

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/response"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	name := c.Params("name")
	fmt.Println("id: ", id)
	fmt.Println("name: ", name)
	var user model.User
	database.DB.Find(&user, id)
	res := getUserResponse{
		ID:           int(user.ID),
		Email:        user.Email,
		Name:         user.Name,
		Lastname:     user.Lastname,
		DepartmentID: user.DepartmentID,
		RoleCode:     user.RoleCode,
	}
	return response.JSON(c, res)
}
