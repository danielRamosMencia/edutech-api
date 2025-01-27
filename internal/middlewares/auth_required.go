package middlewares

import (
	"fmt"
	"log"
	"strings"

	"github.com/danielRamosMencia/edutech-api/internal/constans"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	var stringToken string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		stringToken = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		stringToken = c.Cookies("token")
	}

	if stringToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "middle-err-000",
		})
	}

	byteToken, err := jwt.Parse(stringToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexepected sigining method: %s", jwtToken.Header["alg"])
		}

		return []byte(constans.Envs.JwtSecret), nil
	})
	if err != nil {
		log.Println("Error verifying token: ", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "middle-err-001",
		})
	}

	claims, ok := byteToken.Claims.(jwt.MapClaims)
	if !ok || !byteToken.Valid {
		log.Println("Error verifying token claims: ", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "middle-err-002",
		})
	}

	c.Locals("user_claims", claims)

	return c.Next()
}
