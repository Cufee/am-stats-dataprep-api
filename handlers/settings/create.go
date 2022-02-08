package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/settings/types"
	"github.com/gofiber/fiber/v2"
)

func CreateNewSettings(c *fiber.Ctx) error {
	var response handlers.ResponseJSON

	var settingsData types.GenerationSettings
	if err := c.BodyParser(&settingsData); err != nil {
		response.Error.Message = "Error parsing settings"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	id, err := settings.CreateNewSettings(settingsData)
	if err != nil {
		response.Error.Message = "Error creating settings"
		response.Error.Context = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = id
	return c.JSON(response)
}
