package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/settings"

	"byvko.dev/repo/am-stats-dataprep-api/stats/presets"
	api "github.com/byvko-dev/am-types/api/generic/v1"
	types "github.com/byvko-dev/am-types/dataprep/settings/v1"
	"github.com/gofiber/fiber/v2"
)

func CreateNewSettings(c *fiber.Ctx) error {
	var response api.ResponseWithError

	var settingsData types.GenerationSettings
	if err := c.BodyParser(&settingsData); err != nil {
		response.Error = api.ResponseError{
			Message: "Error parsing settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	settingsData.Options = presets.GetPresetByName(settingsData.Preset)
	id, err := settings.CreateNewSettings(settingsData.OwnerId, settingsData)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error creating settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = id
	return c.JSON(response)
}
