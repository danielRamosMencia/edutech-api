package main

import (
	"fmt"

	"github.com/danielRamosMencia/edutech-api/internal/constants"
	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/danielRamosMencia/edutech-api/internal/routes"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	constants.LoadGlobalEnvs()
	db.ConnectDatabase()
	validations.InitValidator()
	zap_logger.InitLogger()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Edutech API - FIBER",
		AppName:       "Edutech API",
		ErrorHandler:  middlewares.FiberErrorHandler,
		BodyLimit:     10 * 1024 * 1024, // 10 MB
	})

	routes.AppRoutes(app)

	app.Listen(constants.Envs.ServerPort)

	fmt.Println("Hello, World!")
}
