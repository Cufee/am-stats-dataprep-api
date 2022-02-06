package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

// HTTP client
var clientHTTP = &http.Client{Timeout: 10 * time.Second}

// GetJSON - Send a GET request and decode response to target
func GetJSON(url string, target interface{}) error {
	// Send Request
	data, err := clientHTTP.Get(url)
	if err != nil {
		return err
	}
	// Decode
	return json.NewDecoder(data.Body).Decode(target)
}
