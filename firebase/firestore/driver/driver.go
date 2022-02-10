package driver

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type Driver struct {
	client *firestore.Client
}

var globalClient *firestore.Client

func InitFirebaseApp() error {
	app, err := firebase.NewApp(context.Background(), nil) // Auth is pulled from the environment variable "GOOGLE_APPLICATION_CREDENTIALS"
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}
	store, err := app.Firestore(context.Background())
	if err != nil {
		return fmt.Errorf("error initializing firestore: %v", err)
	}
	globalClient = store
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
	doc, err := d.client.Collection(collection).Doc(id).Get(context.Background())
	if err != nil {
		return fmt.Errorf("error getting document: %v", err)
	}
	return doc.DataTo(out)
}

func (d *Driver) ReplaceDocumentByID(collection, id string, data interface{}) error {
	_, err := d.client.Collection(collection).Doc(id).Set(context.Background(), data)
	if err != nil {
		return fmt.Errorf("error setting document: %v", err)
	}
	return nil
}

func (d *Driver) CreateDocumentInCollection(collection string, data interface{}) (string, error) {
	newDoc := d.client.Collection(collection).NewDoc()
	if newDoc == nil {
		return "", fmt.Errorf("error creating new document")
	}
	_, err := newDoc.Set(context.Background(), data)
	if err != nil {
		d.client.Doc(newDoc.ID).Delete(context.Background())
		return "", fmt.Errorf("error setting document: %v", err)
	}
	return newDoc.ID, nil
}

func (d *Driver) UpdateDocumentByID(collection, id string, update map[string]interface{}) error {
	var payload []firestore.Update
	for k, v := range update {
		payload = append(payload, firestore.Update{Path: k, Value: v})
	}
	_, err := d.client.Collection(collection).Doc(id).Update(context.Background(), payload)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}
	return nil

}
