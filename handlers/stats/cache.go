package stats

import (
	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/stats"
	statsapi "byvko.dev/repo/am-stats-dataprep-api/stats-api"
	"byvko.dev/repo/am-stats-dataprep-api/stats/cache"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets"
	"github.com/gofiber/fiber/v2"
)

func CacheStatsFromSettings(c *fiber.Ctx) error {
	var response handlers.ResponseJSON

	settingsID := c.Params("id")
	if settingsID == "" {
		response.Error = &handlers.ResponseError{
			Message: "Missing required parameters",
			Context: "Settings ID is required",
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	userSettings, err := settings.GetSettingsByID(settingsID)
	if err != nil {
		response.Error = &handlers.ResponseError{
			Message: "Error getting settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	if userSettings.Player.ID == 0 || userSettings.Player.Realm == "" {
		response.Error = &handlers.ResponseError{
			Message: "Invalid settings",
			Context: "Player ID and Realm are required",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if !userSettings.UseCustomOptions {
		userSettings.Options = presets.GetPresetByName(userSettings.StylePreset)
		userSettings.Options.Locale = userSettings.Locale
	}

	// Get stats
	statsData, err := statsapi.GetStatsByPlayerID(userSettings.Player.ID, userSettings.Player.Realm, 0)
	if err != nil {
		response.Error = &handlers.ResponseError{
			Message: "Error getting stats",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Check for passed in options -- use default for now
	completeCards, err := stats.CompilePlayerStatsCards(statsData, userSettings.Options)
	if err != nil {
		response.Error = &handlers.ResponseError{
			Message: "Error compiling stats",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	completeCards.StylePreset = userSettings.StylePreset

	// Save to cache
	id, err := cache.CreateStatsCache(completeCards)
	if err != nil {
		response.Error = &handlers.ResponseError{
			Message: "Error creating stats cache",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = id
	return c.JSON(response)
}

func CachedStatsFromID(c *fiber.Ctx) error {
	var response handlers.ResponseJSON

	cacheId := c.Params("id")
	if cacheId == "" {
		response.Error = &handlers.ResponseError{
			Message: "Missing required parameters",
			Context: "Settings ID is required",
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Save to cache
	compiledCards, err := cache.GetStatsCacheByID(cacheId)
	if err != nil {
		response.Error = &handlers.ResponseError{
			Message: "Error getting stats cache",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = compiledCards
	return c.JSON(response)
}
