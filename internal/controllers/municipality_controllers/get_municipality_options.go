package municipality_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/municipality_services"
	"github.com/gofiber/fiber/v2"
)

func GetMunicipalityOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	municipalities, status, message, err := municipality_services.SelectMunicipalityOptions(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "municipality-err-006",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": municipalities,
	})
}
