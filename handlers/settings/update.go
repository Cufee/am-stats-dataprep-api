package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	api "github.com/byvko-dev/am-types/api/v1"
	types "github.com/byvko-dev/am-types/dataprep/v1/settings"
	"github.com/gofiber/fiber/v2"
)

func UpdateSettingsByID(c *fiber.Ctx) error {
	var response api.ResponseWithError
	settingsID := c.Params("id")
	if settingsID == "" {
		response.Error = api.ResponseError{
			Message: "Missing required parameters",
			Context: "Settings ID is required",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var settingsData types.GenerationSettings
	if err := c.BodyParser(&settingsData); err != nil {
		response.Error = api.ResponseError{
			Message: "Error parsing settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err := settings.UpdateSettingsByID(settingsID, settingsData)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error updating settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	settingsData.OwnerId = "" // remove owner id from response
	response.Data = settingsData
	return c.JSON(response)
}
