package grade_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/grade_services"
	"github.com/gofiber/fiber/v2"
)

func GetGrade(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	gradeId := c.Params("id")

	grade, status, message, err := grade_services.SelectGrade(ctx, gradeId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "grade-err-001",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": grade,
	})
}
