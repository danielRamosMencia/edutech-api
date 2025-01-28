package country_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/country_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetCountryRoutes(router fiber.Router) {
	r := router.Group("/countries")
	r.Use(middlewares.AuthRequired)

	r.Get("/", country_controllers.GetCountries)
	r.Get("/:id", country_controllers.GetCountry)
}
