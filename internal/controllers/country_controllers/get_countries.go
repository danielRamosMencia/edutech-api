package country_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constans"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/services/country_services"
	"github.com/gofiber/fiber/v2"
)

func GetCountries(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constans.ContextTimeOut)
	defer cancel()

	pagination := helpers.MapPagination(c)

	countries, status, message, err := country_services.SelectCountries(ctx, pagination)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "country-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message":       message,
		"records":       countries,
		"page":          pagination.Offset/pagination.Limit + 1,
		"page_size":     pagination.Limit,
		"total_records": len(countries),
	})
}
