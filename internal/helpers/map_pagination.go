package helpers

import (
	"github.com/danielRamosMencia/edutech-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

func MapPagination(c *fiber.Ctx) models.PaginationParams {
	page, limit := c.QueryInt("page", 1), c.QueryInt("limit", 50)

	offset := (page - 1) * limit

	return models.PaginationParams{
		Offset: offset,
		Limit:  limit,
	}
}
