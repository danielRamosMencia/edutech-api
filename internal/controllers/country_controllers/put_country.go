package country_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models/country_models"
	"github.com/danielRamosMencia/edutech-api/internal/services/country_services"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PutCountry(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "country-err-003",
		})
	}

	var updateCountry country_models.UpdateCountry
	err = c.BodyParser(&updateCountry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de actualizar país incorrectos",
			"code":  "country-err-003",
		})
	}

	err = validations.Validate.Struct(updateCountry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "country-err-003",
		})
	}

	countryId := c.Params("id")

	status, message, err := country_services.UpdateCountry(ctx, updateCountry, countryId, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "country-err-003",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
