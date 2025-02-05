package employee_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/employee_services"
	"github.com/gofiber/fiber/v2"
)

func GetEmployee(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	employeeId := c.Params("id")

	employee, status, message, err := employee_services.SelectEmployee(ctx, employeeId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "employee-err-001",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": employee,
	})
}
