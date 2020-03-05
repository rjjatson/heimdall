package model

const (
	// SenderIDJSONTag json tag
	SenderIDJSONTag = "sender_id"

	// ReceiverIDJSONTag json tag
	ReceiverIDJSONTag = "receiver_id"
	// MessageIDJSONTag json tag
	MessageIDJSONTag = "id"
	// MessageTypeJSONTag json tag
	MessageTypeJSONTag = "type"
)

// Request is common struct from client on request-response scheme
type Request struct {
	SenderID string `json:"sender_id"`
	ID       string `json:"id"`
	Type     string `json:"type"`
}

// Response is common struct from server on request-response scheme
//
// response ID have the same id as request message
// error message and error code omitted on nil error
type Response struct {
	ReceiverID   string `json:"receiver_id"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	ErrorMessage string `json:"error_message,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty"`
}
