package stats

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats"
	statsapi "byvko.dev/repo/am-stats-dataprep-api/stats-api"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/presets"
	api "github.com/byvko-dev/am-types/api/generic/v1"
	types "github.com/byvko-dev/am-types/stats/v1"
	"github.com/gofiber/fiber/v2"
)

func GenerateStatsWithOptions(c *fiber.Ctx) error {
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
	statsData, err := statsapi.GetStatsFromRequest(request)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error getting stats",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	options := presets.LoadOPresetByName(request.Preset)
	completeCards, err := stats.CompilePlayerStatsCards(statsData, options, request.Locale, request.Style)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error compiling stats",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	completeCards.Style = request.Style

	response.Data = completeCards
	return c.JSON(response)
}

// func GenerateStatsFromSettings(c *fiber.Ctx) error {
// 	var response api.ResponseWithError

// 	settingsID := c.Params("id")
// 	if settingsID == "" {
// 		response.Error = api.ResponseError{
// 			Message: "Missing required parameters",
// 			Context: "Settings ID is required",
// 		}

// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	userSettings, err := settings.GetSettingsByID(settingsID)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error getting settings",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}
// 	if userSettings.Player.ID == 0 || userSettings.Player.Realm == "" {
// 		response.Error = api.ResponseError{
// 			Message: "Invalid settings",
// 			Context: "Player ID and Realm are required",
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	if !userSettings.UseCustomOptions {
// 		userSettings.Options = presets.GetPresetByName(userSettings.Preset)
// 		userSettings.Options.Locale = userSettings.Locale
// 	}

// 	// Get stats
// 	statsData, err := statsapi.GetStatsByPlayerID(userSettings.Player.ID, userSettings.Player.Realm, 0)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error getting stats",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}

// 	// Check for passed in options -- use default for now
// 	completeCards, err := stats.CompilePlayerStatsCards(statsData, userSettings.Options, userSettings.Locale, userSettings.Style)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error compiling stats",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}
// 	completeCards.Style = userSettings.Style

// 	response.Data = completeCards
// 	return c.JSON(response)
// }
