package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/pkg/response"
)

func CheckIsAdmin(c *fiber.Ctx) error {
	roleCode := c.Locals("roleCode")
	if roleCode != "1" {
		return response.Error(c, fiber.StatusNotAcceptable, "You are not admin")
	}
	return c.Next()
}
