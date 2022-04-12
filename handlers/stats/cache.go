package stats

// func CacheStatsFromSettings(c *fiber.Ctx) error {
// 	var response api.ResponseWithError

// 	settingsID := c.Params("id")
// 	if settingsID == "" {
// 		response.Error = api.ResponseError{
// 			Message: "Missing required parameters",
// 			Context: "Settings ID is required",
// 		}

// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	userSettings, err := settings.GetSettingsByID(settingsID)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error getting settings",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}
// 	if userSettings.Player.ID == 0 || userSettings.Player.Realm == "" {
// 		response.Error = api.ResponseError{
// 			Message: "Invalid settings",
// 			Context: "Player ID and Realm are required",
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	if !userSettings.UseCustomOptions {
// 		userSettings.Options = presets.GetPresetByName(userSettings.Preset)
// 		userSettings.Options.Locale = userSettings.Locale
// 	}

// 	// Get stats
// 	statsData, err := statsapi.GetStatsByPlayerID(userSettings.Player.ID, userSettings.Player.Realm, 0)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error getting stats",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}

// 	// Check for passed in options -- use default for now
// 	completeCards, err := stats.CompilePlayerStatsCards(statsData, userSettings.Options, userSettings.Style)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error compiling stats",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}
// 	completeCards.Style = userSettings.Style

// 	// Save to cache
// 	id, err := cache.CreateStatsCache(completeCards)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error creating stats cache",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}

// 	response.Data = id
// 	return c.JSON(response)
// }

// func CacheStatsFromOptions(c *fiber.Ctx) error {
// 	var response api.ResponseWithError

// 	var request types.StatsRequest
// 	if err := c.BodyParser(&request); err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error parsing request",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	if (request.PID == 0) || (request.Realm == "") {
// 		response.Error = api.ResponseError{
// 			Message: "Missing required parameters",
// 			Context: "Player ID and Realm are required",
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	// Get stats
// 	statsData, err := statsapi.GetStatsFromRequest(request)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error getting stats",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}

// 	options := presets.GetPresetByName(request.Preset)
// 	completeCards, err := stats.CompilePlayerStatsCards(statsData, options, request.Style)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error compiling stats",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}
// 	completeCards.Style = request.Style

// 	// Save to cache
// 	id, err := cache.CreateStatsCache(completeCards)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error creating stats cache",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}

// 	response.Data = id
// 	return c.JSON(response)
// }

// func CachedStatsFromID(c *fiber.Ctx) error {
// 	var response api.ResponseWithError

// 	cacheId := c.Params("id")
// 	if cacheId == "" {
// 		response.Error = api.ResponseError{
// 			Message: "Missing required parameters",
// 			Context: "Settings ID is required",
// 		}

// 		return c.Status(fiber.StatusBadRequest).JSON(response)
// 	}

// 	// Save to cache
// 	compiledCards, err := cache.GetStatsCacheByID(cacheId)
// 	if err != nil {
// 		response.Error = api.ResponseError{
// 			Message: "Error getting stats cache",
// 			Context: err.Error(),
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(response)
// 	}

// 	response.Data = compiledCards
// 	return c.JSON(response)
// }
