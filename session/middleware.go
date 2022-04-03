package session

import (
	"github.com/byvko-dev/am-core/logs"
	api "github.com/byvko-dev/am-types/api/v1"
	"github.com/gofiber/fiber/v2"
)

func SessionCheckMiddleware(ctx *fiber.Ctx) error {
	user, err := GetUserFromSession(ctx)
	if err != nil || user.ID == "" {
		logs.Error("Error getting user from session", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(api.ResponseWithError{
			Error: api.ResponseError{
				Message: "Unauthorized",
			},
		})
	}

	return ctx.Next()

}
