package handler

import (
	"heimdall/internal/router"

	"github.com/sirupsen/logrus"
)

// HandleEcho handles echo message
func HandleEcho(msg []byte, rw *router.ResponseWriter) {
	logrus.Debug("executing echo")
}
