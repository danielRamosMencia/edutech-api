package department_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/services/department_services"
	"github.com/gofiber/fiber/v2"
)

func GetDepartments(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	pagination := helpers.MapPagination(c)

	departments, status, message, err := department_services.SelectDepartments(ctx, pagination)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "department-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": departments,
	})
}
