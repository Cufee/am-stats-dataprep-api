package session

import (
	"errors"

	types "github.com/byvko-dev/am-types/users/v1"
	"github.com/gofiber/fiber/v2"
)

func GetUserFromSession(ctx *fiber.Ctx) (types.GlobalUser, error) {
	// session := ctx.Cookies("am_session", strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", ""))
	// if session == "" {
	// 	return types.GlobalUser{}, errors.New("missing session cookie")
	// }

	// sessionDetails, err := sessions.FindSessionDetails(session)
	// if err != nil {
	// 	return types.GlobalUser{}, err
	// }

	// user, err := users.FindUserByID(sessionDetails.UserID)
	// if err != nil {
	// 	return types.GlobalUser{}, err
	// }
	return types.GlobalUser{}, errors.New("not implemented")
}
