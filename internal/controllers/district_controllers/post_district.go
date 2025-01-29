package district_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models/district_models"
	"github.com/danielRamosMencia/edutech-api/internal/services/district_services"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PostDistrict(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "district-err-002",
		})
	}

	var newDistrict district_models.CreateDistrict
	err = c.BodyParser(&newDistrict)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de crear distrito incorrectos",
			"code":  "district-err-002",
		})
	}

	err = validations.Validate.Struct(newDistrict)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "district-err-002",
		})
	}

	status, message, err := district_services.InsertDistrict(ctx, newDistrict, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "district-err-002",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
