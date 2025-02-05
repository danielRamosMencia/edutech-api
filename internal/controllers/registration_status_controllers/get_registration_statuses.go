package registration_status_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/services/registration_status_services"
	"github.com/gofiber/fiber/v2"
)

func GetRegistrationStatuses(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	pagination := helpers.MapPagination(c)

	registrationStatuses, status, message, err := registration_status_services.SelectRegistrationStatuses(ctx, pagination)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "registration-status-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": registrationStatuses,
	})
}
