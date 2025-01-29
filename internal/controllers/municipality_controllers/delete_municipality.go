package municipality_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/municipality_services"
	"github.com/gofiber/fiber/v2"
)

func DeleteMunicipality(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	municipalityId := c.Params("id")

	status, message, err := municipality_services.DeleteMunicipality(ctx, municipalityId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "municipality-err-005",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
