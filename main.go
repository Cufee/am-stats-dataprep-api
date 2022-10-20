package main

import (
	"fmt"
	"os"

	"byvko.dev/repo/am-stats-dataprep-api/handlers/settings"
	"byvko.dev/repo/am-stats-dataprep-api/handlers/stats"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/presets"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	presets.Init() // idk why init functions there do not run otherwise

	// Define routes
	app := fiber.New()

	// Logger
	app.Use(logger.New())

	apiV1 := app.Group("/v1")

	// Open routes
	statsV1 := apiV1.Group("/stats")
	statsV1.Post("/options", stats.GenerateStatsWithOptions)
	// statsV1.Get("/cache/:id", stats.CachedStatsFromID)
	// statsV1.Get("/settings/:id", stats.GenerateStatsFromSettings)
	// statsV1.Post("/options/cache", stats.CacheStatsFromOptions)

	settingsV1 := apiV1.Group("/settings")
	settingsV1.Get("/:id", settings.GetSettingsByID)
	// settingsV1.Get("/:id/cache", stats.CacheStatsFromSettings)
	settingsV1.Post("/:id", settings.UpdateSettingsByID)
	settingsV1.Put("/", settings.CreateNewSettings)

	panic(app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
