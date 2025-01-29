package district_controllers

import (
	"context"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/services/district_services"
	"github.com/gofiber/fiber/v2"
)

func GetDistrict(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeOut)
	defer cancel()

	districtId := c.Params("id")

	district, status, message, err := district_services.SelectDistrict(ctx, districtId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
			"code":    "district-err-001",
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"records": district,
	})
}
