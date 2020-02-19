package hub

import (
	"heimdall/internal/client"
	"log"
)

// New create new hub
func New(inbound chan []byte) *Hub {
	return &Hub{
		inbound: inbound,
	}
	// todo inject inbound channel

}

// Hub connects server and client
type Hub struct {
	clients  map[string]*client.Client
	inbound  chan []byte
	outbound chan interface{}
}

// Run starts hub processing
func (h *Hub) Run() {
	go h.listen()
}

func (h *Hub) listen() {
	for {
		select {
		case msg := <-h.inbound:
			log.Println(string(msg))
		}
	}
}

// AddClient to the hub
func (h *Hub) AddClient(userID string, client *client.Client) {

}

// RemoveClient removes client from client list
func (h *Hub) RemoveClient(userID string) {

}

// GetInbound gets inbound channel of hub
func (h *Hub) GetInbound() chan []byte {
	return h.inbound
}
