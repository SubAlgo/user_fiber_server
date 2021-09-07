package authController

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/password"
	"github.com/subAlgo/userFiber/pkg/response"
	"strings"
	"unicode/utf8"
)

type SignupRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Name            string `json:"name"`
	Lastname        string `json:"lastname"`
	Phone           string `json:"phone"`
}

func Signup(c *fiber.Ctx) error {
	var data SignupRequest
	var err error

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// check email
	data.Email = strings.ToLower(strings.TrimSpace(data.Email))
	if !govalidator.IsEmail(data.Email) {
		return response.Error(c, fiber.StatusBadRequest, "Email invalid")
	}

	// check password
	if n := utf8.RuneCountInString(strings.TrimSpace(data.Password)); n < 8 || n > 24 {
		return response.Error(c, fiber.StatusBadRequest, errPasswordLength)
	}

	// check confirm password
	if strings.Compare(data.Password, data.ConfirmPassword) != 0 {
		return response.Error(c, fiber.StatusBadRequest, errConfirmPassword)
	}

	// check name
	data.Name = strings.TrimSpace(data.Name)
	if n := utf8.RuneCountInString(data.Name); n == 0 {
		return response.Error(c, fiber.StatusBadRequest, errNameRequire)
	}

	// check lastname
	data.Lastname = strings.TrimSpace(data.Lastname)
	if n := utf8.RuneCountInString(data.Lastname); n == 0 {
		return response.Error(c, fiber.StatusBadRequest, errLastnameRequire)
	}

	//
	data.Phone = strings.TrimSpace(data.Phone)

	// hash password
	hashPassword, _ := password.Hash(data.Password)

	user := &model.User{
		Email:        data.Email,
		Password:     hashPassword,
		Name:         data.Name,
		Lastname:     data.Lastname,
		RoleCode:     "2",
		DepartmentID: 1,
		Salary:       0,
	}

	if err = database.DB.Create(&user).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	var u model.User
	database.DB.Where("email = ?", data.Email).First(&u)

	if n := utf8.RuneCountInString(strings.TrimSpace(data.Phone)); n > 0 {
		phone := &model.Phone{
			PhoneNo: data.Phone,
			UserID:  u.ID,
		}
		if err = database.DB.Create(&phone).Error; err != nil {
			return response.Error(c, fiber.StatusBadRequest, "set phone number")
		}
	}

	return response.JsonMessage(c, "Success")
}
