package department_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constans"
	"github.com/danielRamosMencia/edutech-api/internal/services/department_services"
	"github.com/gofiber/fiber/v2"
)

func DeleteDepartment(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constans.ContextTimeOut)
	defer cancel()

	departmentId := c.Params("id")

	status, message, err := department_services.DeleteDepartment(ctx, departmentId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "department-err-005",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
