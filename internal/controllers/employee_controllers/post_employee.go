package employee_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models/employee_models"
	"github.com/danielRamosMencia/edutech-api/internal/services/employee_services"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PostEmployee(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "employee-err-002",
		})
	}

	var newEmployee employee_models.CreateEmployee
	err = c.BodyParser(&newEmployee)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de crear empleado incorrectos",
			"code":  "employee-err-002",
		})
	}

	err = validations.Validate.Struct(newEmployee)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "employee-err-002",
		})
	}

	status, message, err := employee_services.InsertEmployee(ctx, newEmployee, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "employee-err-002",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})

}
