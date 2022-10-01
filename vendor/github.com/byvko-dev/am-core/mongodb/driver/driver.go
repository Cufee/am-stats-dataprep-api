package driver

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	errors "github.com/byvko-dev/am-core/errors/database"

	"github.com/byvko-dev/am-core/helpers/maps"
	"github.com/byvko-dev/am-core/logs"
	dbErrors "github.com/byvko-dev/am-core/mongodb/errors"
	"github.com/byvko-dev/am-core/mongodb/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	db            *mongo.Database
	databaseName  string
	connectionUri string
}

func (c *Client) Reconnect() error {
	return InitGlobalConnetion(c.connectionUri, c.databaseName)
}

type StreamChunk struct {
	Data map[string]interface{}
	Err  error
}

var globalClient Client

func InitGlobalConnetion(connectionUri, databaseName string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionUri))
	if err != nil {
		logs.Error("InitGlobalConnetion / Failed to create client: %v", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		cancel()
		logs.Error("InitGlobalConnetion / Failed to connect to cluster: %v", err)
		return err
	}

	// Cancel the context if the program finishes
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
	}()

	if err := pingDatabase(client); err != nil {
		return err
	}

	globalClient.db = client.Database(databaseName)
	globalClient.connectionUri = connectionUri
	globalClient.databaseName = databaseName
	return nil
}

func pingDatabase(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func iterateChangeStream(routineCtx context.Context, stream *mongo.ChangeStream, out chan StreamChunk) {
	defer stream.Close(routineCtx)
	for stream.Next(routineCtx) {
		data := make(map[string]interface{})
		if err := stream.Decode(&data); err != nil {
			out <- StreamChunk{Err: err}
			return
		}
		out <- StreamChunk{Data: data}
	}
}

func mapToBsonM(data map[string]interface{}) (bson.M, error) {
	if data == nil {
		return bson.M{}, nil
	}
	encoded, err := bson.Marshal(data)
	if err != nil {
		return bson.M{}, nil
	}
	var out bson.M
	err = bson.Unmarshal(encoded, &out)
	if err != nil {
		return bson.M{}, fmt.Errorf("mapToBsonM / Failed to unmarshal: %v", err)
	}
	return out, nil
}

func NewClient() (*Client, error) {
	err := pingDatabase(globalClient.db.Client())
	if err != nil {
		logs.Warning("Failed to ping database: %v", err)
		return &globalClient, globalClient.Reconnect()
	}
	return &globalClient, nil
}

func (c *Client) GetDocumentWithFilter(collection types.Collection, filter map[string]interface{}, out interface{}) error {
	bsonFilter, err := mapToBsonM(filter)
	if err != nil {
		return err
	}
	collectionRef := c.db.Collection(collection.GetName())
	if collectionRef == nil {
		return errors.ErrCollectionDoesNotExist
	}
	result := collectionRef.FindOne(context.TODO(), bsonFilter)
	if result == nil {
		return errors.ErrDocumentNotFound
	}
	return dbErrors.Wrap(result.Decode(out))
}

func (c *Client) InsertDocument(collection types.Collection, data interface{}) (interface{}, error) {
	defer c.validateIndexes(collection)

	collectionRef := c.db.Collection(collection.GetName())
	if collectionRef == nil {
		return nil, errors.ErrCollectionDoesNotExist
	}
	res, err := collectionRef.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, err
}

func (c *Client) UpdateDocumentWithFilter(collection types.Collection, filter map[string]interface{}, update map[string]interface{}, upsert bool) error {
	defer c.validateIndexes(collection)

	bsonFilter, err := mapToBsonM(filter)
	if err != nil {
		return err
	}

	set := make(map[string]interface{})
	if _, ok := update["$set"]; !ok {
		set["$set"] = update
	} else {
		set = update
	}
	bsonUpdate, err := mapToBsonM(set)
	if err != nil {
		return err
	}

	opts := options.UpdateOptions{}
	opts.SetUpsert(upsert)

	collectionRef := c.db.Collection(collection.GetName())
	if collectionRef == nil {
		return errors.ErrCollectionDoesNotExist
	}
	_, err = collectionRef.UpdateOne(context.TODO(), bsonFilter, bsonUpdate, &opts)
	return dbErrors.Wrap(err)
}

func (c *Client) SubscribeWithFilter(collection types.Collection, filter map[string]interface{}, project map[string]interface{}) (chan StreamChunk, context.CancelFunc, error) {
	// Pipeline runs on the event document, not on the document itself :facepalm:
	// We need to fix the passed in filter to match the event document
	filterFixed := make(map[string]interface{})
	maps.Flatten("fullDocument", filter, filterFixed)
	bsonFilter, err := mapToBsonM(filterFixed)
	if err != nil {
		return nil, nil, err
	}

	var pipeline []bson.M
	pipeline = append(pipeline, bson.M{"$match": bsonFilter})

	collectionRef := c.db.Collection(collection.GetName())
	if collectionRef == nil {
		return nil, nil, errors.ErrCollectionDoesNotExist
	}

	opts := options.ChangeStream()
	opts.SetFullDocument(options.UpdateLookup)
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := collectionRef.Watch(ctx, pipeline, opts)
	if err != nil {
		cancel()
		return nil, nil, dbErrors.Wrap(err)
	}

	chunkChan := make(chan StreamChunk)
	go iterateChangeStream(ctx, stream, chunkChan)
	return chunkChan, cancel, nil
}

func (c *Client) Aggregate(collection types.Collection, pipeline []bson.D) ([]bson.D, error) {
	collectionRef := c.db.Collection(collection.GetName())
	cursor, err := collectionRef.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	var result []bson.D
	if err := cursor.All(context.TODO(), &result); err != nil {
		return nil, dbErrors.Wrap(err)
	}
	return result, nil
}
