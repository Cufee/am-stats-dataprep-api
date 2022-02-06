package main

import (
	"fmt"
	"log"
	"os"

	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"byvko.dev/repo/am-stats-dataprep-api/localization"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"golang.org/x/text/language"
)

func main() {
	localization.InitLocalizer(language.Russian.String())

	// Define routes
	app := fiber.New()

	// Logger
	app.Use(logger.New())
	// CORS
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	apiV1 := app.Group("/api/v1")

	apiV1.Post("/stats", handlers.GenerateStatsWithOptions)

	log.Print(app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
