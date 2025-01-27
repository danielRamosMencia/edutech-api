package country_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constans"
	"github.com/danielRamosMencia/edutech-api/internal/services/country_services"
	"github.com/gofiber/fiber/v2"
)

func GetCountry(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constans.ContextTimeOut)
	defer cancel()

	countryId := c.Params("id")

	country, status, message, err := country_services.SelectCountry(ctx, countryId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "country-err-001",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": country,
	})
}
