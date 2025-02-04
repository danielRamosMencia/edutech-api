package signature_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/signature_services"
	"github.com/gofiber/fiber/v2"
)

func GetSignatureOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	signatures, status, message, err := signature_services.SelectSignatureOptions(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "signature-err-006",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": signatures,
	})
}
