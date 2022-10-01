package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/byvko-dev/am-core/logs"
)

type Options struct {
	Timeout time.Duration
}

var defaultOptions = Options{
	Timeout: time.Second * 10,
}

func Send(url string, method string, headers map[string]string, payload []byte, target interface{}, options ...Options) (int, error) {
	opts := defaultOptions
	if len(options) > 0 {
		opts.Timeout = options[0].Timeout
	}
	var err error
	var bodyBytes []byte
	var resp *http.Response
	defer func() {
		// Logging
		if err != nil || (resp != nil && resp.StatusCode < 200 || resp.StatusCode > 299) {
			logs.Error("URL: %v", url)
			logs.Error("Headers: %v", headers)
			logs.Error("Payload: %v", string(payload))
			logs.Error("Response: %v", string(bodyBytes))
			logs.Error("Error: %v", err)
		}
	}()

	// Prep request
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(payload))
	if err != nil {
		return 0, err
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Set payload headers
	if payload != nil {
		req.Header.Set("content-type", "application/json")
	}

	// Send request
	client := &http.Client{Timeout: opts.Timeout}
	resp, err = client.Do(req)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return 0, err
	}
	if target != nil {
		// Read body
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, err
		}
		err = json.Unmarshal(bodyBytes, target)
	}
	return resp.StatusCode, err
}
