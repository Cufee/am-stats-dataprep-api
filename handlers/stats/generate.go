package stats

import (
	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/stats"
	statsapi "byvko.dev/repo/am-stats-dataprep-api/stats-api"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/gofiber/fiber/v2"
)

func GenerateStatsWithOptions(c *fiber.Ctx) error {
	var response handlers.ResponseJSON

	var request types.BasicStatsRequest
	if err := c.BodyParser(&request); err != nil {
		response.Error.Message = "Error parsing request body"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if (request.PlayerID == 0) || (request.Realm == "") {
		response.Error.Message = "Missing required parameters"
		response.Error.Context = "Player ID and Realm are required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Get stats
	statsData, err := statsapi.GetStatsByPlayerID(request.PlayerID, request.Realm, 0)
	if err != nil {
		response.Error.Message = "Error getting stats"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Check for passed in options -- use default for now
	options := presets.GetPresetByName(request.Profile)
	options.Locale = request.Locale
	completeCards, err := stats.CompilePlayerStatsCards(statsData, options)
	if err != nil {
		response.Error.Message = "Error compiling stats"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	completeCards.StylePreset = request.Profile

	response.Data = completeCards
	return c.JSON(response)
}

func GenerateStatsFromSettings(c *fiber.Ctx) error {
	var response handlers.ResponseJSON

	settingsID := c.Params("id")
	if settingsID == "" {
		response.Error.Message = "Missing required parameters"
		response.Error.Context = "Settings ID is required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	userSettings, err := settings.GetSettingsByID(settingsID)
	if err != nil {
		response.Error.Message = "Error getting settings"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Get stats
	statsData, err := statsapi.GetStatsByPlayerID(userSettings.Player.ID, userSettings.Player.Realm, 0)
	if err != nil {
		response.Error.Message = "Error getting stats"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Check for passed in options -- use default for now
	completeCards, err := stats.CompilePlayerStatsCards(statsData, userSettings.Options)
	if err != nil {
		response.Error.Message = "Error compiling stats"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	completeCards.StylePreset = userSettings.StylePreset

	response.Data = completeCards
	return c.JSON(response)
}
