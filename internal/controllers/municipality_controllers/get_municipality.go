package municipality_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/municipality_services"
	"github.com/gofiber/fiber/v2"
)

func GetMunicipality(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	municipalityId := c.Params("id")

	municipality, status, message, err := municipality_services.SelectMunicipality(ctx, municipalityId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "municipality-err-001",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": municipality,
	})
}
