package main

import (
	"fmt"
	"heimdall/internal/server"
)

func main() {
	fmt.Println("heimdall service")
	srv := server.New()
	srv.Serve()
}
