package handler

import (
	"encoding/json"
	"heimdall/internal/router"
	"heimdall/pkg/model"

	"github.com/sirupsen/logrus"
)

// HandleEcho handles echo message
func HandleEcho(msg []byte, rw *router.ResponseWriter) {
	logrus.Debug("executing echo")

	var echoRequest model.EchoRequest
	json.Unmarshal(msg, &echoRequest) // todo handle error

	rw.WriteResponse(model.EchoResponse{
		Message: echoRequest.Message,
	})
}
