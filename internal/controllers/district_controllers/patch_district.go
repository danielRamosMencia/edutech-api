package district_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/services/district_services"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PatchDistrict(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "district-err-004",
		})
	}

	var request models.SwicthActive
	err = c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de actualizar distrito incorrectos",
			"code":  "district-err-004",
		})
	}

	value, message, err := validations.ActiveRequired(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": message,
			"code":  "district-err-004",
		})
	}

	districtId := c.Params("id")

	status, message, err := district_services.ActiveDistrict(ctx, value, districtId, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "district-err-004",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
