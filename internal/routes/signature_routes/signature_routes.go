package signature_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/signature_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetSignatureRoutes(router fiber.Router) {

	r := router.Group("/signatures")
	r.Use(middlewares.AuthRequired)

	r.Get("/", signature_controllers.GetSignatures)
	r.Get("/options", signature_controllers.GetSignatureOptions)
	r.Get("/:id", signature_controllers.GetSignature)
	r.Post("/", signature_controllers.PostSignature)
	r.Put("/:id", signature_controllers.PutSignature)
	r.Patch("/:id", signature_controllers.PatchSignature)
	r.Delete("/:id", signature_controllers.DeleteSignature)
}
