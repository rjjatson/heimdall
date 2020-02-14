package main

import (
	"fmt"
	"heimdall/internal/server"
)

func main() {
	fmt.Println("heimdall service")
	server.Serve()
}
