package country_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constans"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models/country_models"
	"github.com/danielRamosMencia/edutech-api/internal/services/country_services"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PostCountry(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constans.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "country-err-002",
		})
	}

	var newCountry country_models.CreateCoutry
	err = c.BodyParser(&newCountry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de crear país incorrectos",
			"code":  "country-err-002",
		})
	}

	err = validations.Validate.Struct(newCountry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "country-err-002",
		})
	}

	status, message, err := country_services.InsertCountry(ctx, newCountry, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "country-err-002",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
