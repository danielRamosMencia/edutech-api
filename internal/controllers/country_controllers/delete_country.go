package country_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constans"
	"github.com/danielRamosMencia/edutech-api/internal/services/country_services"
	"github.com/gofiber/fiber/v2"
)

func DeleteCountry(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constans.ContextTimeOut)
	defer cancel()

	countryId := c.Params("id")

	status, message, err := country_services.DeleteCountry(ctx, countryId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "country-err-005",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
