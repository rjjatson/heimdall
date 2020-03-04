package hub

import (
	"heimdall/internal/client"
	"heimdall/internal/router"

	"github.com/sirupsen/logrus"
)

// New create new hub
func New(router *router.Router, inbound chan []byte, outbound chan []byte) *Hub {
	return &Hub{
		router:   router,
		inbound:  inbound,
		outbound: outbound,
		clients:  make(map[string]*client.Client),
	}
	// todo inject inbound channel

}

// Hub connects server and client
type Hub struct {
	router   *router.Router
	clients  map[string]*client.Client
	inbound  chan []byte
	outbound chan []byte
}

// Run starts hub processing
func (h *Hub) Run() {
	go h.listen()
}

func (h *Hub) listen() {
	for {
		select {
		case msg := <-h.inbound:
			logrus.WithFields(
				logrus.Fields{"package": "hub", "message": string(msg)}).
				Debug("incoming message")
			h.router.Route(msg)
		}
	}
}

// AddClient to the hub
func (h *Hub) AddClient(c *client.Client) {
	h.clients[c.GetUserID()] = c
}

// RemoveClient removes client from client list
func (h *Hub) RemoveClient(userID string) {

}

// GetInbound gets inbound channel of hub
func (h *Hub) GetInbound() chan []byte {
	return h.inbound
}
