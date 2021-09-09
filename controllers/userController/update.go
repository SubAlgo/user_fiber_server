package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/response"
	"strings"
	"unicode/utf8"
)

type updateUserRequest struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func Update(c *fiber.Ctx) error {
	var userModel model.User
	var userReq updateUserRequest

	userID := c.Locals("userID")

	database.DB.Find(&userModel, userID)
	if userModel.Email == "" {
		return response.Error(c, fiber.StatusInternalServerError, "not available")
	}
	// parse data
	if err := c.BodyParser(&userReq); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	// check data
	userReq.Name = strings.TrimSpace(userReq.Name)
	if n := utf8.RuneCountInString(userReq.Name); n <= 0 {
		return response.Error(c, fiber.StatusBadRequest, "name require")
	}
	userReq.Lastname = strings.TrimSpace(userReq.Lastname)
	if n := utf8.RuneCountInString(userReq.Lastname); n <= 0 {
		return response.Error(c, fiber.StatusBadRequest, "lastname require")
	}

	userModel.Name = userReq.Name
	userModel.Lastname = userReq.Lastname
	if err := database.DB.Save(&userModel).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.JsonMessage(c, "Update data success!!!")
}
