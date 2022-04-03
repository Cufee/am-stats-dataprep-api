package init

import (
	"github.com/byvko-dev/am-core/helpers"
	mongodb "github.com/byvko-dev/am-core/mongodb/driver"
)

func Init() {
	// Initialize the mongodb connection
	mongoUri, _ := helpers.MustGetEnv("MONGO_URI")[0].(string)
	databaseName, _ := helpers.MustGetEnv("MONGO_DATABASE")[0].(string)
	err := mongodb.InitGlobalConnetion(mongoUri, databaseName)
	if err != nil {
		panic(err)
	}

}
