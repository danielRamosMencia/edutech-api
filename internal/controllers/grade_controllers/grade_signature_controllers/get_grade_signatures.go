package grade_signature_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/grade_services/grade_signature_services"
	"github.com/gofiber/fiber/v2"
)

func GetGradeSignatures(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	gradeId := c.Params("id")

	gradeSignatures, status, message, err := grade_signature_services.SelectGradeSignatures(ctx, gradeId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "grade-signatures-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": gradeSignatures,
	})

}
