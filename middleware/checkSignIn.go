package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/subAlgo/userFiber/database"
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/response"
	"github.com/subAlgo/userFiber/pkg/token"
)

// CheckSignIn check user signIn or not
func CheckSignIn(c *fiber.Ctx) error {
	cookieToken := c.Cookies("jwt")
	// fmt.Println(cookieToken)
	tk, err := jwt.ParseWithClaims(cookieToken, &jwt.StandardClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(token.SecretKey), nil
	})
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "unauthenticated")
	}

	claims := tk.Claims.(*jwt.StandardClaims)
	var user model.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.ID == 0 {
		return response.Error(c, fiber.StatusUnauthorized, "unauthenticated")
	}
	// save userID in ctx
	c.Locals("userID", user.ID)
	c.Locals("roleCode", user.RoleCode)
	return c.Next()
}
