package grade_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/grade_services"
	"github.com/gofiber/fiber/v2"
)

func GetGradeOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	grades, status, message, err := grade_services.SelectGradeOptions(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "grade-err-006",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": grades,
	})
}
