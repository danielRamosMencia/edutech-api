package district_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/services/district_services"
	"github.com/gofiber/fiber/v2"
)

func GetDistricts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	pagination := helpers.MapPagination(c)

	districts, status, message, err := district_services.SelectDistricts(ctx, pagination)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "district-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": districts,
	})
}
