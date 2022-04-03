package settings

import (
	"byvko.dev/repo/am-stats-dataprep-api/session"
	"byvko.dev/repo/am-stats-dataprep-api/settings"
	"byvko.dev/repo/am-stats-dataprep-api/settings/types"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets"
	"github.com/byvko-dev/am-core/logs"
	api "github.com/byvko-dev/am-types/api/v1"
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

	user, err := session.GetUserFromSession(c)
	if err != nil || user.ID == "" {
		logs.Error("Error getting user from session", err)
		return c.Status(fiber.StatusUnauthorized).JSON(api.ResponseWithError{
			Error: api.ResponseError{
				Message: "Unauthorized",
			},
		})
	}

	settingsData.Options = presets.GetPresetByName(settingsData.StylePreset)
	id, err := settings.CreateNewSettings(user.ID, settingsData)
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

// func CreateNewSettingsWithID(c *fiber.Ctx) error {
// 	var response handlers.ResponseJSON

// 	id := c.Params("id")
// 	if id == "" {
// 		response.Error = &handlers.ResponseError{
// 			Message: "Missing required parameters",
// 			Context: "ID is required",
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	var settingsData types.GenerationSettings
// 	if err := c.BodyParser(&settingsData); err != nil {
// 		response.Error = &handlers.ResponseError{
// 			Message: "Error parsing settings",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	user, err := session.GetUserFromSession(c)
// 	if err != nil || user.ID == "" {
// 		logs.Error("Error getting user from session", err)
// 		return c.Status(fiber.StatusUnauthorized).JSON(handlers.ResponseJSON{
// 			Error: &handlers.ResponseError{
// 				Message: "Unauthorized",
// 			},
// 		})
// 	}

// 	logs.Debug("%+v", user)
// 	logs.Debug("%+v", settingsData)

// 	settingsData.Options = presets.GetPresetByName(settingsData.StylePreset)
// 	err = settings.CreateNewSettingsWithID(id, user.ID, settingsData)
// 	if err != nil {
// 		response.Error = &handlers.ResponseError{
// 			Message: "Error creating settings",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}

// 	response.Data = id
// 	return c.JSON(response)
// }
