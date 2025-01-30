package grade_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/grade_services"
	"github.com/gofiber/fiber/v2"
)

func DeleteGrade(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	gradeId := c.Params("id")

	status, message, err := grade_services.DeleteGrade(ctx, gradeId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "grade-err-005",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
