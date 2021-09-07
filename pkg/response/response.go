package response

import "github.com/gofiber/fiber/v2"

// JSON return not modify response
func JSON(c *fiber.Ctx, i interface{}) error {
	return c.JSON(i)
}

// JsonMessage set return message to JSON format.
func JsonMessage(c *fiber.Ctx, message string) error {
	m := Msg{Message: message}
	return c.JSON(m)
}

// Error set error response
func Error(c *fiber.Ctx, errCode int, errMessage string) error {
	c.Status(errCode)
	return c.JSON(fiber.Map{
		"message": errMessage,
	})
}
