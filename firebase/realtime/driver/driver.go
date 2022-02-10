package driver

import (
	"context"
	"fmt"
	"path"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

type Driver struct {
	client *db.Client
}

var globalClient *db.Client

func InitFirebaseApp() error {
	app, err := firebase.NewApp(context.Background(), &databaseConfig) // Auth is pulled from the environment variable "GOOGLE_APPLICATION_CREDENTIALS"
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}
	realtime, err := app.Database(context.Background())
	if err != nil {
		return fmt.Errorf("error initializing firebase: %v", err)
	}
	globalClient = realtime
	return nil
}

func NewDriver() (*Driver, error) {
	if globalClient == nil {
		err := InitFirebaseApp()
		if err != nil {
			return nil, err
		}
	}
	return &Driver{
		client: globalClient,
	}, nil
}

func (d *Driver) GetDocumentByID(collection, id string, out interface{}) error {
	ref := d.client.NewRef(path.Join(collection, id))
	if ref == nil {
		return fmt.Errorf("error creating new ref")
	}

	return ref.Get(context.Background(), out)
}

func (d *Driver) ReplaceDocumentByID(collection, id string, data interface{}) error {
	ref := d.client.NewRef(path.Join(collection, id))
	if ref == nil {
		return fmt.Errorf("error creating new ref")
	}
	return ref.Set(context.Background(), data)
}

func (d *Driver) CreateDocumentInCollection(collection string, data interface{}) (string, error) {
	ref := d.client.NewRef(collection)
	if ref == nil {
		return "", fmt.Errorf("error creating new ref")
	}

	updateRef, err := ref.Push(context.Background(), data)
	if err != nil {
		return "", fmt.Errorf("error creating new document: %v", err)
	}

	return updateRef.Key, nil
}

func (d *Driver) UpdateDocumentByID(collection, id string, update map[string]interface{}) error {
	ref := d.client.NewRef(path.Join(collection, id))
	if ref == nil {
		return fmt.Errorf("error creating new ref")
	}

	return ref.Update(context.Background(), update)

}
