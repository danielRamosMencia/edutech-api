package helpers

import (
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetClaims(c *fiber.Ctx) (models.SessionData, error) {
	claims, ok := c.Locals("user_claims").(jwt.MapClaims)
	if !ok {
		return models.SessionData{}, fiber.ErrUnauthorized
	}

	return models.SessionData{
		Id:       claims["id"].(string),
		Username: claims["user_name"].(string),
		Email:    claims["email"].(string),
		Active:   claims["active"].(bool),
		RoleId:   claims["role_id"].(string),
		Role:     claims["role"].(string),
	}, nil
}
