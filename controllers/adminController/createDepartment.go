package adminController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/response"
	"strings"
	"unicode/utf8"
)

type createDepartmentRequest struct {
	Title string `json:"title"`
}

func CreateDepartment(c *fiber.Ctx) error {
	var err error
	var department model.Department
	//var req createDepartmentRequest

	if err = c.BodyParser(&department); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if n := utf8.RuneCountInString(strings.TrimSpace(department.Title)); n <= 0 {
		return response.Error(c, fiber.StatusBadRequest, "department name require")
	}

	if err = database.DB.Create(&department).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.JsonMessage(c, "Create new department success")
}
