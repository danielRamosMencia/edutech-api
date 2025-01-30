package grade_routes

import (
	"github.com/danielRamosMencia/edutech-api/internal/controllers/grade_controllers"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetGradeRoutes(router fiber.Router) {
	r := router.Group("/grades")
	r.Use(middlewares.AuthRequired)

	r.Get("/", grade_controllers.GetGrades)
	r.Get("/options", grade_controllers.GetGradeOptions)
	r.Get("/:id", grade_controllers.GetGrade)
	r.Post("/", grade_controllers.PostGrade)
	r.Put("/:id", grade_controllers.PutGrade)
	r.Patch("/:id", grade_controllers.PatchGrade)
	r.Delete("/:id", grade_controllers.DeleteGrade)
}
