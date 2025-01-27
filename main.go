package main

import (
	"fmt"

	"github.com/danielRamosMencia/edutech-api/internal/constans"
	"github.com/danielRamosMencia/edutech-api/internal/db"
	"github.com/danielRamosMencia/edutech-api/internal/middlewares"
	"github.com/danielRamosMencia/edutech-api/internal/routes"
	"github.com/danielRamosMencia/edutech-api/internal/validations"
	"github.com/danielRamosMencia/edutech-api/internal/zap_logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	constans.LoadGlobalEnvs()
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

	app.Listen(constans.Envs.ServerPort)

	fmt.Println("Hello, World!")
}
