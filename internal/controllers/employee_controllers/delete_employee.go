package employee_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/employee_services"
	"github.com/gofiber/fiber/v2"
)

func DeleteEmployee(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	employeeId := c.Params("id")

	status, message, err := employee_services.DeleteEmployee(ctx, employeeId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "employee-err-005",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
