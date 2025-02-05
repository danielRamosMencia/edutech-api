package employee_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/services/employee_services"
	"github.com/gofiber/fiber/v2"
)

func GetEmployees(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	pagination := helpers.MapPagination(c)

	employees, status, message, err := employee_services.SelectEmployees(ctx, pagination)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "employee-err-000",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": employees,
	})
}
