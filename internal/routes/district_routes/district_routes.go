package district_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/district_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetDistrictRoutes(router fiber.Router) {
	r := router.Group("/districts")
	r.Use(middlewares.AuthRequired)

	r.Get("/", district_controllers.GetDistricts)
	r.Get("/options", district_controllers.GetDistrictOptions)
	r.Get("/:id", district_controllers.GetDistrict)
	r.Post("/", district_controllers.PostDistrict)
	r.Put("/:id", district_controllers.PutDistrict)
	r.Patch("/:id", district_controllers.PatchDistrict)
	r.Delete("/:id", district_controllers.DeleteDistrict)
}
