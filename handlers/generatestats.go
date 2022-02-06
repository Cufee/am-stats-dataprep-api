package handlers

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats"
	statsapi "byvko.dev/repo/am-stats-dataprep-api/stats-api"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
	"github.com/gofiber/fiber/v2"
)

func GenerateStatsWithOptions(c *fiber.Ctx) error {
	var response ResponseJSON

	var request types.StatsRequest
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
	options := presets.GetPresetByName(request.SettingsID)
	options.Locale = request.Locale
	completeCards, err := stats.CompilePlayerStatsCards(statsData, options)
	if err != nil {
		response.Error.Message = "Error compiling stats"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	completeCards.StyleProfile = request.SettingsID

	response.Data = completeCards
	return c.JSON(response)
}
