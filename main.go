package main

import (
	"fmt"
	"os"

	database "byvko.dev/repo/am-stats-dataprep-api/database/init"
	"byvko.dev/repo/am-stats-dataprep-api/handlers/settings"
	"byvko.dev/repo/am-stats-dataprep-api/handlers/stats"
	"github.com/byvko-dev/am-core/helpers"
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
	origins, _ := helpers.MustGetEnv("CORS_ALLOW_ORIGINS")[0].(string)
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowOrigins:     origins,
		AllowCredentials: true,
	}))

	apiV1 := app.Group("/v1")

	openV1 := apiV1.Group("/open")
	// secureV1 := apiV1.Group("/secure", session.SessionCheckMiddleware)
	secureV1 := apiV1.Group("/secure")

	// Open routes
	statsOpenV1 := openV1.Group("/stats")
	statsOpenV1.Post("/options", stats.GenerateStatsWithOptions)
	statsOpenV1.Get("/cache/:id", stats.CachedStatsFromID)
	statsOpenV1.Get("/settings/:id", stats.GenerateStatsFromSettings)

	settingsOpenV1 := openV1.Group("/settings")
	settingsOpenV1.Get("/:id", settings.GetSettingsByID)

	// Secure routes
	settingsSecureV1 := secureV1.Group("/settings")
	settingsSecureV1.Put("/", settings.CreateNewSettings)
	settingsSecureV1.Post("/:id", settings.UpdateSettingsByID)

	// Internal -- needs to be moved to a separate service or hidden somehow
	internalV1 := apiV1.Group("/internal")
	statsInternalV1 := internalV1.Group("/stats")
	statsInternalV1.Post("/options/cache", stats.CacheStatsFromOptions)
	statsInternalV1.Get("/settings/:id/cache", stats.CacheStatsFromSettings)

	// Other init logic
	go database.Init()

	panic(app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
