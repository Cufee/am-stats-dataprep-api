package errors

import (
	"errors"

	"github.com/byvko-dev/am-core/errors/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func Wrap(err error) error {
	if errors.Is(mongo.ErrNoDocuments, err) {
		return database.ErrDocumentNotFound
	}
	return err
}
