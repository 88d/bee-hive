package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Client is an middleman between the websocket connection and the hub.
type Client struct {
	WebSocket    *websocket.Conn
	Send         chan *Message
	SessionToken string
	UserID       string
}

func NewClient(ws *websocket.Conn, userId string, sessionToken string) *Client {
	return &Client{
		Send:         make(chan *Message, 64),
		WebSocket:    ws,
		UserID:       userId,
		SessionToken: sessionToken,
	}
}

// readPump pumps messages from the websocket connection to the hub.
func (c *Client) readPump() {
	defer func() {
		root.UnRegister(c)
		c.WebSocket.Close()
	}()
	c.WebSocket.SetReadLimit(maxMessageSize)
	c.WebSocket.SetReadDeadline(time.Now().Add(pongWait))
	c.WebSocket.SetPongHandler(func(string) error {
		c.WebSocket.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		var message Message
		if err := c.WebSocket.ReadJSON(&message); err != nil {
			return
		}
		message.Author = c.UserID
		if root.MessageReceived != nil {
			root.MessageReceived(&message, c, root)
		}
		//root.Broadcast(&message)
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.WebSocket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.WriteMessage(websocket.CloseMessage)
				return
			}
			if err := c.WriteJSON(message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.WriteMessage(websocket.PingMessage); err != nil {
				return
			}
		}
	}
}

func (c *Client) WriteJSON(msg *Message) error {
	c.WebSocket.SetWriteDeadline(time.Now().Add(writeWait))
	return c.WebSocket.WriteJSON(msg)
}

func (c *Client) WriteMessage(msgType int) error {
	c.WebSocket.SetWriteDeadline(time.Now().Add(writeWait))
	return c.WebSocket.WriteMessage(msgType, []byte{})
}
