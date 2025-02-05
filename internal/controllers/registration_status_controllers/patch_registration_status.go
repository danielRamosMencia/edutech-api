package registration_status_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/danielRamosMencia/edutech-api/internal/services/registration_status_services"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PatchRegistrationStatus(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "registration-status-err-004",
		})
	}

	var request models.SwicthActive
	err = c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de actualizar registro incorrectos",
			"code":  "registration-status-err-004",
		})
	}

	value, message, err := validations.ActiveRequired(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": message,
			"code":  "registration-status-err-004",
		})
	}

	regStatusId := c.Params("id")

	status, message, err := registration_status_services.ActiveRegistrationStatus(ctx, value, regStatusId, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "registration-status-err-004",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
