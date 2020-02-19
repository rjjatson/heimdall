package client

import (
	"log"

	"github.com/gorilla/websocket"
)

// New create new client
func New(c *websocket.Conn, inbound chan []byte) *Client {
	return &Client{
		connection: c,
		inbound:    inbound,
	}
}

const (
	maxMessageSize = 512
)

// Client manages client communication
type Client struct {
	connection *websocket.Conn
	inbound    chan []byte
}

// Run starts listen and write process from client
func (c *Client) Run() {
	go c.listen()
	go c.write()
}

func (c *Client) listen() {
	defer func() {
		c.connection.Close()
	}()
	c.connection.SetReadLimit(maxMessageSize)

	// todo: create ping-pong mechanism
	for {
		_, msg, err := c.connection.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		c.inbound <- msg
	}
}

func (c *Client) write() {

}
