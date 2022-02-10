package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/settings/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets"
	"github.com/gofiber/fiber/v2"
)

func CreateNewSettings(c *fiber.Ctx) error {
	var response handlers.ResponseJSON

	var settingsData types.GenerationSettings
	if err := c.BodyParser(&settingsData); err != nil {
		response.Error = &handlers.ResponseError{
			Message: "Error parsing settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	settingsData.Options = presets.GetPresetByName(settingsData.StylePreset)
	id, err := settings.CreateNewSettings(settingsData)
	if err != nil {
		response.Error = &handlers.ResponseError{
			Message: "Error creating settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = id
	return c.JSON(response)
}
