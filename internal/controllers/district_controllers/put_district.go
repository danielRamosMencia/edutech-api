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

func PutDistrict(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "district-err-003",
		})
	}

	var updateDistrict district_models.UpdateDistrict
	err = c.BodyParser(&updateDistrict)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de actualizar distrito incorrectos",
			"code":  "district-err-003",
		})
	}

	err = validations.Validate.Struct(updateDistrict)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "district-err-003",
		})
	}

	districtId := c.Params("id")

	status, message, err := district_services.UpdateDistrict(ctx, updateDistrict, districtId, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "district-err-004",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
