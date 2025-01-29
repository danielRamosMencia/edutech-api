package municipality_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models/municipality_models"
	"github.com/danielRamosMencia/edutech-api/internal/services/municipality_services"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PostMunicipality(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "municipality-err-002",
		})
	}

	var newMunicipality municipality_models.CreateMunicipality
	err = c.BodyParser(&newMunicipality)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de crear municipio incorrectos",
			"code":  "municipality-err-002",
		})
	}

	err = validations.Validate.Struct(newMunicipality)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "municipality-err-002",
		})
	}

	status, message, err := municipality_services.InsertMunicipality(ctx, newMunicipality, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "municipality-err-002",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
