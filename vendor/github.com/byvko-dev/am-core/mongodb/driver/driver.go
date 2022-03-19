package driver

import (
	"context"
	"fmt"
	"time"

	"github.com/byvko-dev/am-core/helpers"
	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-core/mongodb/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	*mongo.Database
	databaseName  string
	connectionUri string
}

func (c *Client) Reconnect() error {
	return InitGlobalConnetion(c.connectionUri, c.databaseName)
}

type streamChunk struct {
	data interface{}
	err  error
}

var globalClient *Client

func InitGlobalConnetion(connectionUri, databaseName string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionUri))
	if err != nil {
		logs.Error("InitGlobalConnetion / Failed to create client: %v", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logs.Error("InitGlobalConnetion / Failed to connect to cluster: %v", err)
		return err
	}

	if err := pingDatabase(client); err != nil {
		return err
	}

	globalClient.Database = client.Database(databaseName)
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

func completeMongoURL() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/?tls=false", helpers.MustGetEnv("MONGO_USER", "MONGO_PASSWORD", "MONGO_HOST", "MONGO_PORT")...)
}

func iterateChangeStream(routineCtx context.Context, stream *mongo.ChangeStream, out chan streamChunk) {
	defer stream.Close(routineCtx)
	for stream.Next(routineCtx) {
		logs.Debug("iterateChangeStream / Received chunk: %v", stream.Current)
		var data string
		if err := stream.Decode(&data); err != nil {
			out <- streamChunk{err: err}
			return
		}
		out <- streamChunk{data: data}
	}
}

func mapToBson(data map[string]interface{}) (bson.D, error) {
	if data == nil {
		return bson.D{}, nil
	}

	encoded, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}
	var bsonDoc bson.D
	if err := bson.Unmarshal(encoded, &bsonDoc); err != nil {
		return nil, err
	}
	return bsonDoc, nil
}

func NewClient() (*Client, error) {
	err := pingDatabase(globalClient.Client())
	if err != nil {
		logs.Warning("Failed to ping database: %v", err)
		return globalClient, globalClient.Reconnect()
	}
	return globalClient, nil
}

func (c *Client) GetDocumentWithFilter(collection types.Collection, filter map[string]interface{}, out interface{}) error {
	bsonFilter, err := mapToBson(filter)
	if err != nil {
		return err
	}
	collectionRef := c.Collection(collection.GetName())
	cursor, err := collectionRef.Find(context.TODO(), bsonFilter)
	if err != nil {
		return err
	}
	if err := cursor.Decode(out); err != nil {
		return err
	}
	return nil
}

func (c *Client) InsertDocument(collection types.Collection, data interface{}) (interface{}, error) {
	defer c.validateIndexes(collection)

	collectionRef := c.Collection(collection.GetName())
	res, err := collectionRef.InsertOne(context.TODO(), data)
	return res.InsertedID, err
}

func (c *Client) UpdateDocumentWithFilter(collection types.Collection, filter map[string]interface{}, update map[string]interface{}, upsert bool) error {
	defer c.validateIndexes(collection)

	bsonFilter, err := mapToBson(filter)
	if err != nil {
		return err
	}

	var bsonUpdate bson.D
	set := make(map[string]interface{})
	if _, ok := update["$set"]; !ok {
		set["$set"] = update
	} else {
		set = update
	}
	bsonUpdate, err = mapToBson(set)
	if err != nil {
		return err
	}

	opts := options.UpdateOptions{}
	opts.SetUpsert(upsert)

	collectionRef := c.Collection(collection.GetName())
	_, err = collectionRef.UpdateOne(context.TODO(), bsonFilter, bsonUpdate, &opts)
	return err
}

func (c *Client) SubscribeWithFilter(collection types.Collection, filter map[string]interface{}, project map[string]interface{}, out chan interface{}) error {
	filterMap := make(map[string]interface{})
	filterMap["$match"] = filter
	bsonFilter, err := mapToBson(filterMap)
	if err != nil {
		return err
	}

	var pipeline mongo.Pipeline
	projectMap := make(map[string]interface{})
	if project != nil {
		projectMap["$project"] = project
		bsonProject, err := mapToBson(projectMap)
		if err != nil {
			return err
		}
		pipeline = append(pipeline, bsonProject)
	}
	pipeline = append(pipeline, bsonFilter)

	logs.Debug("SubscribeWithFilter / Pipeline: %#v", pipeline)

	collectionRef := c.Collection(collection.GetName())
	stream, err := collectionRef.Watch(context.TODO(), pipeline)
	if err != nil {
		return err
	}

	routineCtx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	chunkChan := make(chan streamChunk)
	go iterateChangeStream(routineCtx, stream, chunkChan)

	logs.Debug("SubscribeWithFilter / Starting to listen for changes")
	defer func() {
		logs.Debug("SubscribeWithFilter / Closing stream")
	}()

	// Get chunks in channel until error
	for {
		select {
		case chunk := <-chunkChan:
			logs.Debug("SubscribeWithFilter / Received chunk: %v", chunk)
			if chunk.err != nil {
				cancelFn()
				return chunk.err
			}
			out <- chunk.data
		case <-routineCtx.Done():
			return nil
		}
	}
}
