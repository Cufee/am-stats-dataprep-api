package main

import (
	"fmt"
	"os"

	"byvko.dev/repo/am-stats-dataprep-api/handlers/settings"
	"byvko.dev/repo/am-stats-dataprep-api/handlers/stats"
	"byvko.dev/repo/am-stats-dataprep-api/session"
	"github.com/byvko-dev/am-core/helpers"
	"github.com/byvko-dev/am-core/mongodb/driver"
	shims "github.com/byvko-dev/am-core/shims/database/mongodb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize the database connection
	mongoUri, _ := helpers.MustGetEnv("MONGO_URI")[0].(string)
	databaseName, err := shims.GetDatabaseNameFromURI(mongoUri)
	if err != nil {
		panic(err)
	}
	err = driver.InitGlobalConnetion(mongoUri, databaseName)
	if err != nil {
		panic(err)
	}

	// // test setup
	// c := driver.NewClient()
	// testDoc := map[string]interface{}{
	// 	"test": "test",
	// }

	// collection := types.NewCollection("test", nil)
	// c.UpdateDocumentWithFilter(collection, nil, testDoc, true)

	// out := make(chan interface{})
	// err := c.SubscribeWithFilter(collection, testDoc, nil, out)
	// logs.Error(err)

	// go func() {
	// 	for {
	// 		fmt.Println(<-out)
	// 	}
	// }()

	// Define routes
	app := fiber.New()

	// Logger
	app.Use(logger.New())
	// CORS
	origins, _ := helpers.MustGetEnv("CORS_ALLOW_ORIGINS")[0].(string)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowOrigins:     origins,
	}))

	apiV1 := app.Group("/v1")

	statsV1 := apiV1.Group("/stats")
	statsV1.Post("/", stats.GenerateStatsWithOptions)
	statsV1.Get("/settings/:id", stats.GenerateStatsFromSettings)
	statsV1.Get("/settings/:id/cache", stats.CacheStatsFromSettings)
	statsV1.Get("/cache/:id", stats.CachedStatsFromID)

	settingsV1 := apiV1.Group("/settings")
	settingsV1.Get("/:id", settings.GetSettingsByID)
	settingsV1.Put("/", session.SessionCheckMiddleware, settings.CreateNewSettings)
	settingsV1.Put("/:id", session.SessionCheckMiddleware, settings.CreateNewSettingsWithID)
	settingsV1.Post("/:id", session.SessionCheckMiddleware, settings.UpdateSettingsByID)

	// Catch all
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect(os.Getenv("FRONTEND_URL"))
	})

	panic(app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
