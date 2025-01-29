package department_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/department_services"
	"github.com/gofiber/fiber/v2"
)

func GetDepartmentOptions(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	departments, status, message, err := department_services.SelectDepartmentOptions(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "department-err-006",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": departments,
	})
}
