package department_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/helpers"
	"github.com/danielRamosMencia/edutech-api/internal/models/department_models"
	"github.com/danielRamosMencia/edutech-api/internal/services/department_services"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/gofiber/fiber/v2"
)

func PutDepartment(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	sessionData, err := helpers.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Inautorizado",
			"code":  "department-err-003",
		})
	}

	var updateDepartment department_models.UpdateDeparment
	err = c.BodyParser(&updateDepartment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Campos para solicitud de actualizar departamento incorrectos",
			"code":  "department-err-003",
		})
	}

	err = validations.Validate.Struct(updateDepartment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validations.MapValidatorErrors(err),
			"code":  "department-err-003",
		})
	}

	status, message, err := department_services.UpdateDepartment(ctx, c.Params("id"), updateDepartment, sessionData.Username)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"error": message,
			"code":  "department-err-003",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
