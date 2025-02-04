package grade_signature_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/grade_services/grade_signature_services"
	"github.com/gofiber/fiber/v2"
)

func DeleteAssignedSignature(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	gradeId := c.Params("id")
	recordId := c.Params("recordId")

	status, message, err := grade_signature_services.UnassignSignature(ctx, recordId, gradeId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "grade-signature-err-003",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
