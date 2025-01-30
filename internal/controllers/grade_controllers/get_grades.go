package grade_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/services/grade_services"
	"github.com/gofiber/fiber/v2"
)

func GetGrades(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	pagination := helpers.MapPagination(c)

	grades, status, message, err := grade_services.SelectGrades(ctx, pagination)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "grade-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": grades,
	})

}
