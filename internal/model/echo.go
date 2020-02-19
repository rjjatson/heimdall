package model

// EchoRequest request for echo from client
type EchoRequest struct {
	Request
	Message string `type:"message"`
}

// EchoResponse response for echo from server
type EchoResponse struct {
	Response
	Message string `type:"message"`
}
