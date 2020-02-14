package server

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var addr = flag.String("addr", ":8080", "http service address")

// Serve start webservice
func Serve() {
	flag.Parse()
	http.HandleFunc("/connect", connect)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func connect(w http.ResponseWriter, r *http.Request) {
	log.Print("connecting")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte("OK"))
	if err != nil {
		log.Println(err)
		return
	}
}
