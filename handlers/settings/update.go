package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/settings/types"
	"github.com/gofiber/fiber/v2"
)

func UpdateSettingsByID(c *fiber.Ctx) error {
	var response handlers.ResponseJSON

	settingsID := c.Params("id")
	if settingsID == "" {
		response.Error.Message = "Missing required parameters"
		response.Error.Context = "Settings ID is required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var settingsData types.GenerationSettings
	if err := c.BodyParser(&settingsData); err != nil {
		response.Error.Message = "Error parsing settings"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err := settings.UpdateSettingsByID(settingsID, settingsData)
	if err != nil {
		response.Error.Message = "Error updating settings"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = settingsData
	return c.JSON(response)
}
