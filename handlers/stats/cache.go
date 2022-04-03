package stats

import (
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/stats"
	statsapi "byvko.dev/repo/am-stats-dataprep-api/stats-api"
	"byvko.dev/repo/am-stats-dataprep-api/stats/cache"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets"
	"github.com/byvko-dev/am-core/logs"
	api "github.com/byvko-dev/am-types/api/v1"
	types "github.com/byvko-dev/am-types/stats/v1"
	"github.com/gofiber/fiber/v2"
)

func CacheStatsFromSettings(c *fiber.Ctx) error {
	var response api.ResponseWithError

	settingsID := c.Params("id")
	if settingsID == "" {
		response.Error = api.ResponseError{
			Message: "Missing required parameters",
			Context: "Settings ID is required",
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	userSettings, err := settings.GetSettingsByID(settingsID)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error getting settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	if userSettings.Player.ID == 0 || userSettings.Player.Realm == "" {
		response.Error = api.ResponseError{
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
		response.Error = api.ResponseError{
			Message: "Error getting stats",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Check for passed in options -- use default for now
	completeCards, err := stats.CompilePlayerStatsCards(statsData, userSettings.Options)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error compiling stats",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	completeCards.StylePreset = userSettings.StylePreset

	// Save to cache
	id, err := cache.CreateStatsCache(completeCards)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error creating stats cache",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = id
	return c.JSON(response)
}

func CacheStatsFromOptions(c *fiber.Ctx) error {
	var response api.ResponseWithError

	var request types.StatsRequest
	if err := c.BodyParser(&request); err != nil {
		response.Error = api.ResponseError{
			Message: "Error parsing request",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if (request.PID == 0) || (request.Realm == "") {
		response.Error = api.ResponseError{
			Message: "Missing required parameters",
			Context: "Player ID and Realm are required",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Get stats
	statsData, err := statsapi.GetStatsByPlayerID(request.PID, request.Realm, 0)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error getting stats",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	options := presets.GetPresetByName(request.Profile)
	completeCards, err := stats.CompilePlayerStatsCards(statsData, options)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error compiling stats",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	completeCards.StylePreset = request.Profile

	logs.Debug("%+v", completeCards.Cards[1].Rows)

	// Save to cache
	id, err := cache.CreateStatsCache(completeCards)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error creating stats cache",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = id
	return c.JSON(response)
}

func CachedStatsFromID(c *fiber.Ctx) error {
	var response api.ResponseWithError

	cacheId := c.Params("id")
	if cacheId == "" {
		response.Error = api.ResponseError{
			Message: "Missing required parameters",
			Context: "Settings ID is required",
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Save to cache
	compiledCards, err := cache.GetStatsCacheByID(cacheId)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error getting stats cache",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = compiledCards
	return c.JSON(response)
}
