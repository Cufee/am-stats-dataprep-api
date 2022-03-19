package mongodb

import (
	"fmt"
	"net/url"
)

func GetDatabaseNameFromURI(uri string) (string, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	if parsed.Path == "" || len(parsed.Path) < 2 {
		return "", fmt.Errorf("invalid database uri: %s", uri)
	}
	return parsed.Path[1:], nil
}
