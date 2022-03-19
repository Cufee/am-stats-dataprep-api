package users

import (
	"github.com/byvko-dev/am-core/firebase/firestore/driver"
	types "github.com/byvko-dev/am-types/users/v1"
)

const usersCollection = "globalUsers"

func FindUserByID(id string) (types.GlobalUser, error) {
	var user types.GlobalUser

	client, err := driver.NewDriver()
	if err != nil {
		return user, err
	}

	err = client.GetDocumentByID(usersCollection, id, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
