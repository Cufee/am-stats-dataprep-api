package main

import (
	"fmt"
	"os"

	database "byvko.dev/repo/am-stats-dataprep-api/database/init"
	"byvko.dev/repo/am-stats-dataprep-api/handlers/settings"
	"byvko.dev/repo/am-stats-dataprep-api/handlers/stats"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Define routes
	app := fiber.New()

	// Logger
	app.Use(logger.New())

	apiV1 := app.Group("/v1")

	// Open routes
	statsV1 := apiV1.Group("/stats")
	statsV1.Post("/options", stats.GenerateStatsWithOptions)
	statsV1.Get("/cache/:id", stats.CachedStatsFromID)
	statsV1.Get("/settings/:id", stats.GenerateStatsFromSettings)
	statsV1.Post("/options/cache", stats.CacheStatsFromOptions)

	settingsV1 := apiV1.Group("/settings")
	settingsV1.Get("/:id", settings.GetSettingsByID)
	settingsV1.Get("/:id/cache", stats.CacheStatsFromSettings)
	settingsV1.Post("/:id", settings.UpdateSettingsByID)
	settingsV1.Put("/", settings.CreateNewSettings)

	// Other init logic
	go database.Init()

	panic(app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
