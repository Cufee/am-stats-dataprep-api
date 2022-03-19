package session

import (
	"byvko.dev/repo/am-stats-dataprep-api/handlers"
	"github.com/byvko-dev/am-core/logs"
	"github.com/gofiber/fiber/v2"
)

func SessionCheckMiddleware(ctx *fiber.Ctx) error {
	user, err := GetUserFromSession(ctx)
	if err != nil || user.ID == "" {
		logs.Error("Error getting user from session", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(handlers.ResponseJSON{
			Error: &handlers.ResponseError{
				Message: "Unauthorized",
			},
		})
	}

	return ctx.Next()

}
