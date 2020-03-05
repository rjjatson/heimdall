package server

import (
	"flag"
	"heimdall/internal/client"
	"heimdall/internal/hub"
	"heimdall/internal/router"
	"heimdall/pkg/handler"
	"log"
	"net/http"

	"github.com/google/uuid"
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

	// todo cyclic dep?
	router.Add("authorize", hub.AuthorizeUser)

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

	// todo : add authorization
	uid := "guest-" + uuid.New().String()

	c := client.New(uid, conn, s.hub.GetInbound())
	c.Run()
	s.hub.AddClient(c.GetUserID(), c)

	err = conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"connect_notif","user_id":"`+uid+`"}`))
	if err != nil {
		log.Println(err)
		return
	}
}
