package authController

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/password"
	"github.com/subAlgo/userFiber/pkg/response"
	"github.com/subAlgo/userFiber/pkg/token"
	"time"
)

type SignInStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *fiber.Ctx) error {
	var req SignInStruct
	var err error

	if err = c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	var user model.User

	// check user email
	database.DB.Where("email = ?", req.Email).First(&user)
	if user.ID == 0 {
		return response.Error(c, fiber.StatusNotFound, "user not found")
	}

	// check password
	if err = password.Compare(user.Password, req.Password); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "incorrect password")
	}

	// create new token
	var tk string
	tk, err = token.New(user.ID)
	if err != nil {
		fmt.Println(err)
		return response.Error(c, fiber.StatusInternalServerError, "could not login")
	}

	// set fiber cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tk,
		Expires:  time.Now().Add(time.Hour * 2),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return response.JsonMessage(c, tk)
}
