package registration_status_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models/registration_status_models"
	"github.com/danielRamosMencia/edutech-api/internal/services/registration_status_services"
	"github.com/gofiber/fiber/v2"
)

func UpdateRegistrationStatus(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "registration-status-err-003",
		})
	}

	var updateRegStatus registration_status_models.UpdateRegistrationStatus

	err = c.BodyParser(&updateRegStatus)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de actualizar registro incorrectos",
			"code":  "registration-status-err-003",
		})
	}

	regStatusId := c.Params("id")

	status, message, err := registration_status_services.UpdateRegistrationStatus(ctx, updateRegStatus, regStatusId, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "registration-status-err-003",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
