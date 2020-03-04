package server

import (
	"flag"
	"heimdall/internal/client"
	"heimdall/internal/handler"
	"heimdall/internal/hub"
	"heimdall/internal/router"
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
	outbound := make(chan []byte, 1)
	router := router.New(outbound)
	router.Add("echo", handler.HandleEcho) // todo : move to separated file

	hub := hub.New(router, inbound, outbound)
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

	c := client.New("123", conn, s.hub.GetInbound()) // todo remove hardcode userID
	c.Run()
	s.hub.AddClient(c)

	err = conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"connectNotif","id":0}`))
	if err != nil {
		log.Println(err)
		return
	}
}
