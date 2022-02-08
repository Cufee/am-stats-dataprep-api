package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"github.com/gofiber/fiber/v2"
)

func GetSettingsByID(c *fiber.Ctx) error {
	var response handlers.ResponseJSON

	settingsID := c.Params("id")
	if settingsID == "" {
		response.Error.Message = "Missing required parameters"
		response.Error.Context = "Settings ID is required"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	settingsData, err := settings.GetSettingsByID(settingsID)
	if err != nil {
		response.Error.Message = "Error getting settings"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = settingsData
	return c.JSON(response)
}
