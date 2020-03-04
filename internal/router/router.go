package router

import (
	"encoding/json"
	"heimdall/internal/model"

	"github.com/sirupsen/logrus"
)

// Router specifies routes for incoming message
type Router struct {
	handlers   map[string]func([]byte, *ResponseWriter)
	respWriter *ResponseWriter
	outbound   chan []byte
}

// New creates new router
func New(outbound chan []byte) *Router {
	return &Router{
		outbound: outbound,
		handlers: make(map[string]func([]byte, *ResponseWriter)),
	}
}

// Route routes message payload to a matching handler
func (r *Router) Route(msgPayload []byte) {
	var requestSign model.Request
	err := json.Unmarshal(msgPayload, &requestSign)
	if err != nil {
		logrus.Error("unable to unmarshal request ", err)
		return
	}
	rw := NewResponseWriter(requestSign.SenderID, r.outbound)

	if r.handlers[requestSign.Type] == nil {
		logrus.WithFields(
			logrus.Fields{"type": requestSign.Type}).
			Error("handler not found")
		return
	}
	logrus.WithFields(
		logrus.Fields{"type": requestSign.Type, "sender_id": requestSign.SenderID}).
		Debug("routing request")
	r.handlers[requestSign.Type](msgPayload, rw)
}

// Add add function handler to router
func (r *Router) Add(msgType string, handler func([]byte, *ResponseWriter)) {
	// todo add register checking
	r.handlers[msgType] = handler
}

// ResponseWriter writes to client
type ResponseWriter struct {
	receiverID string
	outbound   chan []byte
}

// NewResponseWriter create response writer for sending message
func NewResponseWriter(receiverID string, outbound chan []byte) *ResponseWriter {
	return &ResponseWriter{
		receiverID: receiverID,
	}
}

// WriteResponse write response to corresponding client of response writer
func (resp *ResponseWriter) WriteResponse(msg interface{}) {
	// convert interface to byte
	// send to hub's outbound
	b, err := json.Marshal(msg)
	if err != nil {
		logrus.Error("error marshalling response ", err)
	}
	go func() {
		resp.outbound <- b
	}()
}

// WriteError write error response to corresponding client of response writer
func (resp *ResponseWriter) WriteError() {

}
