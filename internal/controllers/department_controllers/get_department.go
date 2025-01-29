package department_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/department_services"
	"github.com/gofiber/fiber/v2"
)

func GetDepartment(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	deparmentId := c.Params("id")

	department, status, message, err := department_services.SelectDepartment(ctx, deparmentId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "department-err-001",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": department,
	})
}
