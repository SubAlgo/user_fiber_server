package Middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/pkg/response"
)

func CheckContentType(c *fiber.Ctx) error {
	if h := c.Get("Content-Type"); h != "application/json" {
		return response.Error(c, fiber.StatusUnsupportedMediaType, "Unsupported MediaType")
	}
	// Go to next middleware:
	return c.Next()
}
