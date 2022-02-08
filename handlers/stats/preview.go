package stats

import (
	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/stats"
	statsapi "byvko.dev/repo/am-stats-dataprep-api/stats-api"
	"github.com/gofiber/fiber/v2"
)

func PreviewSettings(c *fiber.Ctx) error {
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
	statsData, err := statsapi.GetMockStats()
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
