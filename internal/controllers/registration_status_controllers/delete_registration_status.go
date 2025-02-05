package registration_status_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/registration_status_services"
	"github.com/gofiber/fiber/v2"
)

func DeleteRegistrationStatus(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	regStatusId := c.Params("id")

	status, message, err := registration_status_services.DeleteRegistrationStatus(ctx, regStatusId)
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
