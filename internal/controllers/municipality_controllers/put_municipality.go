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

func PutMunicipality(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "municipality-err-003",
		})
	}

	var updateMunicipality municipality_models.UpdateMunicipality
	err = c.BodyParser(&updateMunicipality)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de actualizar municipio incorrectos",
			"code":  "municipality-err-003",
		})
	}

	err = validations.Validate.Struct(updateMunicipality)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "municipality-err-003",
		})
	}

	municipalityId := c.Params("id")

	status, message, err := municipality_services.UpdateMunicipality(ctx, municipalityId, updateMunicipality, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "municipality-err-003",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
