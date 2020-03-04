package main

import (
	"fmt"
	"heimdall/internal/server"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("heimdall service")
	logrus.SetLevel(logrus.DebugLevel)
	srv := server.New()
	srv.Serve()
}
