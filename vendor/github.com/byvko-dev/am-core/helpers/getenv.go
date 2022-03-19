package helpers

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func MustGetEnv(keys ...string) []interface{} {
	var env []interface{}
	for _, key := range keys {
		if val := os.Getenv(key); val != "" {
			env = append(env, val)
			continue
		}
		panic(fmt.Sprintf("%s is not set", key))
	}
	return env
}
