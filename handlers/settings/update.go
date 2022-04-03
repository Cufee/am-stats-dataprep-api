package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/session"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/settings/types"
	"github.com/byvko-dev/am-core/logs"
	api "github.com/byvko-dev/am-types/api/v1"
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

	user, err := session.GetUserFromSession(c)
	if err != nil || user.ID == "" {
		logs.Error("Error getting user from session", err)
		return c.Status(fiber.StatusUnauthorized).JSON(api.ResponseWithError{
			Error: api.ResponseError{
				Message: "Unauthorized",
			},
		})
	}

	if user.ID != settingsData.OwnerId {
		response.Error = api.ResponseError{
			Message: "Unauthorized",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	err = settings.UpdateSettingsByID(settingsID, settingsData)
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
