package authController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/pkg/response"
	"time"
)

func SignOut(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return response.JsonMessage(c, "success")
}
