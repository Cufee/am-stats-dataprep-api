package main

import (
	"fmt"
	"log"
	"os"

	"byvko.dev/repo/am-stats-dataprep-api/handlers/settings"
	"byvko.dev/repo/am-stats-dataprep-api/handlers/stats"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
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

	statsV1 := apiV1.Group("/stats")
	statsV1.Post("/", stats.GenerateStatsWithOptions)
	statsV1.Get("/settings/:id/preview", stats.PreviewSettings)
	statsV1.Get("/settings/:id", stats.GenerateStatsFromSettings)

	settingsV1 := apiV1.Group("/settings")
	settingsV1.Put("/", settings.CreateNewSettings)
	settingsV1.Get("/:id", settings.GetSettingsByID)
	settingsV1.Post("/:id", settings.UpdateSettingsByID)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
