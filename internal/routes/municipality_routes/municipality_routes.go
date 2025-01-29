package municipality_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/municipality_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetMunicipalityRoutes(router fiber.Router) {
	r := router.Group("/municipalities")
	r.Use(middlewares.AuthRequired)

	r.Get("/", municipality_controllers.GetMunicipalities)
	r.Get("/options", municipality_controllers.GetMunicipalityOptions)
	r.Get("/:id", municipality_controllers.GetMunicipality)
	r.Post("/", municipality_controllers.PostMunicipality)
	r.Put("/:id", municipality_controllers.PutMunicipality)
	r.Patch("/:id", municipality_controllers.PatchMunicipality)
	r.Delete("/:id", municipality_controllers.DeleteMunicipality)
}
