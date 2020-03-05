package hub

import (
	"encoding/json"
	"fmt"
	"heimdall/internal/client"
	"heimdall/internal/router"
	"heimdall/pkg/model"
	"sync"

	"github.com/sirupsen/logrus"
)

// New create new hub
func New(router *router.Router, inbound chan []byte, outbound chan []byte) *Hub {
	return &Hub{
		router:   router,
		inbound:  inbound,
		outbound: outbound,
	}
}

// Hub connects server and client
type Hub struct {
	router   *router.Router
	clients  sync.Map
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
		case msg := <-h.outbound:
			logrus.WithFields(
				logrus.Fields{"package": "hub", "message": string(msg)}).
				Debug("outgoing message")
			h.Send(msg)
		}
	}
}

// Send send message to client
func (h *Hub) Send(msg []byte) {
	var resp model.Response
	json.Unmarshal(msg, &resp)

	c, err := h.GetClient(resp.ReceiverID) // todo change to syncmap
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"package": "hub", "client_id": resp.ReceiverID}).
			Error(err)
		return
	}
	c.SendMessage(msg)
}

// GetClient obtain client with selected user ID
func (h *Hub) GetClient(id string) (*client.Client, error) {
	c, ok := h.clients.Load(id)
	if !ok {
		return nil, fmt.Errorf("client not found")
	}
	return c.(*client.Client), nil
}

// AddClient to the hub
func (h *Hub) AddClient(id string, c *client.Client) {
	h.clients.Store(id, c)
}

// RemoveClient removes client from client list
func (h *Hub) RemoveClient(id string) {
	h.clients.Delete(id)
}

// GetInbound gets inbound channel of hub
func (h *Hub) GetInbound() chan []byte {
	return h.inbound
}
