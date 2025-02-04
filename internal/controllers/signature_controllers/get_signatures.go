package signature_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/services/signature_services"
	"github.com/gofiber/fiber/v2"
)

func GetSignatures(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	pagination := helpers.MapPagination(c)

	signatures, status, message, err := signature_services.SelectSignatures(ctx, pagination)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "signature-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": signatures,
	})
}
