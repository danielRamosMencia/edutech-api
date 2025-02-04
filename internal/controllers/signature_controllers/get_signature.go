package signature_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/signature_services"
	"github.com/gofiber/fiber/v2"
)

func GetSignature(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	signatureId := c.Params("id")

	signature, status, message, err := signature_services.SelectSignature(ctx, signatureId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "signature-err-001",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": signature,
	})
}
