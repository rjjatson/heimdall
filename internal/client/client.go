package client

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// New create new client
func New(userID string, c *websocket.Conn, inbound chan []byte) *Client {
	return &Client{
		userID:     userID,
		connection: c,
		inbound:    inbound,
		outbound:   make(chan []byte, 0),
	}
}

const (
	maxMessageSize = 512

	writeTimeOut = 10 * time.Second
)

// Client manages client communication
type Client struct {
	connection *websocket.Conn
	userID     string
	inbound    chan []byte
	outbound   chan []byte
}

// SendMessage asynchronously send messge to outbound channel
func (c *Client) SendMessage(msg []byte) {
	go func() {
		c.outbound <- msg
	}()
}

// GetUserID get client asscoiated userID
func (c *Client) GetUserID() string {
	return c.userID
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

	// todo: create ping pong mechanism
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
	defer func() {
		c.connection.Close()
	}()

	// todo: create ping pong mechanism
	for {
		select {
		case msg, ok := <-c.outbound:
			c.connection.SetWriteDeadline(time.Now().Add(writeTimeOut))
			if !ok {
				c.connection.WriteMessage(websocket.CloseMessage, []byte{})
			}
			w, err := c.connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)
			if err = w.Close(); err != nil {
				return
			}
		}
	}
}
