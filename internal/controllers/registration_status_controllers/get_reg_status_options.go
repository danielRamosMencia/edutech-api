package registration_status_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/registration_status_services"
	"github.com/gofiber/fiber/v2"
)

func GetRegStatusOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	registrationStatuses, status, message, err := registration_status_services.SelectRegStatusOptions(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "registration-status-err-006",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": registrationStatuses,
	})
}
