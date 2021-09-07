package authController

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Signin(c *fiber.Ctx) error {
	fmt.Println("signin")
	return nil
}
