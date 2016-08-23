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

// Conn is an middleman between the websocket connection and the hub.
type Conn struct {
	WebSocket    *websocket.Conn
	Send         chan *Message
	SessionToken string
	UserID       string
}

func NewConn(ws *websocket.Conn, userId string, sessionToken string) *Conn {
	return &Conn{
		Send:         make(chan *Message, 64),
		WebSocket:    ws,
		UserID:       userId,
		SessionToken: sessionToken,
	}
}

// readPump pumps messages from the websocket connection to the hub.
func (c *Conn) readPump() {
	defer func() {
		h.UnRegister(c)
		c.WebSocket.Close()
	}()
	c.WebSocket.SetReadLimit(maxMessageSize)
	c.WebSocket.SetReadDeadline(time.Now().Add(pongWait))
	c.WebSocket.SetPongHandler(func(string) error { c.WebSocket.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var message Message
		if err := c.WebSocket.ReadJSON(&message); err != nil {
			return
		}
		message.Author = c.UserID
		h.Broadcast(&message)
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Conn) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.WebSocket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.WebSocket.SetWriteDeadline(time.Now().Add(writeWait))
				c.WebSocket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.WebSocket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.WebSocket.WriteJSON(message); err != nil {
				return
			}

		case <-ticker.C:
			c.WebSocket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.WebSocket.WriteJSON(NewPingMessage()); err != nil {
				return
			}
		}
	}
}
