package country_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constans"
	"github.com/danielRamosMencia/edutech-api/internal/services/country_services"
	"github.com/gofiber/fiber/v2"
)

func GetCountryOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constans.ContextTimeOut)
	defer cancel()

	countries, status, message, err := country_services.SelectCountryOptions(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "country-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": countries,
	})
}
