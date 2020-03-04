package model

// EchoRequest request for echo from client
type EchoRequest struct {
	Version string `json:"version"`
	Message string `json:"message"`
}

// EchoResponse response for echo from server
type EchoResponse struct {
	Message string `json:"message"`
}
