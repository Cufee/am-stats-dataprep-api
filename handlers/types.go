package handlers

type ResponseJSON struct {
	Data  interface{}   `json:"data"`
	Error ResponseError `json:"error,omitempty"`
}

type ResponseError struct {
	Message string `json:"message"`
	Context string `json:"context,omitempty"`
}
