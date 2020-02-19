package server

import (
	"flag"
	"heimdall/internal/client"
	"heimdall/internal/hub"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // todo : remove debug
}

var addr = flag.String("addr", ":8080", "http service address")

// New create new server
func New() *Server {
	inbound := make(chan []byte, 1)
	hub := hub.New(inbound)
	hub.Run()
	return &Server{
		hub: hub,
	}
}

// Server serve heidmall service
type Server struct {
	hub *hub.Hub
}

// Serve start webservice
func (s *Server) Serve() {
	flag.Parse()
	http.HandleFunc("/connect", s.connect)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Server) connect(w http.ResponseWriter, r *http.Request) {
	log.Print("connecting")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := client.New(conn, s.hub.GetInbound())
	c.Run()
	s.hub.AddClient("123", c) // todo remove hardcode

	err = conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"connectNotif","id":0}`))
	if err != nil {
		log.Println(err)
		return
	}
}
