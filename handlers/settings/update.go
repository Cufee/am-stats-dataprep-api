package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"byvko.dev/repo/am-stats-dataprep-api/session"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/settings/types"
	"github.com/byvko-dev/am-core/logs"
	"github.com/gofiber/fiber/v2"
)

func UpdateSettingsByID(c *fiber.Ctx) error {
	var response handlers.ResponseJSON

	settingsID := c.Params("id")
	if settingsID == "" {
		response.Error = &handlers.ResponseError{
			Message: "Missing required parameters",
			Context: "Settings ID is required",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var settingsData types.GenerationSettings
	if err := c.BodyParser(&settingsData); err != nil {
		response.Error = &handlers.ResponseError{
			Message: "Error parsing settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := session.GetUserFromSession(c)
	if err != nil || user.ID == "" {
		logs.Error("Error getting user from session", err)
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.ResponseJSON{
			Error: &handlers.ResponseError{
				Message: "Unauthorized",
			},
		})
	}

	err = settings.UpdateSettingsByID(user.ID, settingsID, settingsData)
	if err != nil {
		response.Error = &handlers.ResponseError{
			Message: "Error updating settings",
			Context: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = settingsData
	return c.JSON(response)
}
