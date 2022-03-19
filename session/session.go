package session

import (
	"errors"

	"byvko.dev/repo/am-stats-dataprep-api/database/sessions"
	"byvko.dev/repo/am-stats-dataprep-api/database/users"
	types "github.com/byvko-dev/am-types/users/v1"
	"github.com/gofiber/fiber/v2"
)

func GetUserFromSession(ctx *fiber.Ctx) (types.GlobalUser, error) {
	session := ctx.Cookies("am_session")
	if session == "" {
		return types.GlobalUser{}, errors.New("missing session cookie")
	}

	sessionDetails, err := sessions.FindSessionDetails(session)
	if err != nil {
		return types.GlobalUser{}, err
	}

	user, err := users.FindUserByID(sessionDetails.UserID)
	if err != nil {
		return types.GlobalUser{}, err
	}
	return user, nil
}
