package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	api "github.com/byvko-dev/am-types/api/v1"
	"github.com/gofiber/fiber/v2"
)

func GetSettingsByID(c *fiber.Ctx) error {
	var response api.ResponseWithError

	settingsID := c.Params("id")
	if settingsID == "" {
		response.Error = api.ResponseError{
			Message: "Missing required parameters",
			Context: "Settings ID is required",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	settingsData, err := settings.GetSettingsByID(settingsID)
	if err != nil {
		response.Error = api.ResponseError{
			Message: "Error getting settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	settingsData.OwnerId = "" // remove owner id from response
	response.Data = settingsData
	return c.JSON(response)
}
