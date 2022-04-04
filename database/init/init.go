package init

import (
	"github.com/byvko-dev/am-core/helpers/env"
	"github.com/byvko-dev/am-core/logs"
	mongodb "github.com/byvko-dev/am-core/mongodb/driver"
)

func Init() {
	// Initialize the mongodb connection
	mongoUri := env.MustGet("MONGO_URI")[0].(string)
	databaseName := env.MustGet("MONGO_DATABASE")[0].(string)
	err := mongodb.InitGlobalConnetion(mongoUri, databaseName)
	if err != nil {
		panic(err)
	}
	logs.Info("MongoDB connection initialized")
}
