package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func EcoText(c *fiber.Ctx) error {
	fmt.Println("eco middleware")
	return c.Next()
}
