package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/response"
)

type responseGetUsers struct {
	List []getUserResponse `json:"list"`
}

func GetUsers(c *fiber.Ctx) error {
	var users []model.User
	database.DB.Find(&users)

	var res responseGetUsers
	for _, u := range users {
		x := getUserResponse{
			ID:           int(u.ID),
			Email:        u.Email,
			Name:         u.Name,
			Lastname:     u.Lastname,
			DepartmentID: u.DepartmentID,
			RoleCode:     u.RoleCode,
		}
		res.List = append(res.List, x)
	}
	return response.JSON(c, res)
}
