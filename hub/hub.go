package hub

import "log"

type Hub struct {
	// Registered connections.
	connections map[*Conn]bool

	// Inbound messages from the connections.
	broadcast chan *Message

	// Register requests from the connections.
	register chan *Conn

	// Unregister requests from connections.
	unregister chan *Conn
}

func NewHub() *Hub {
	return &Hub{
		broadcast:   make(chan *Message),
		register:    make(chan *Conn),
		unregister:  make(chan *Conn),
		connections: make(map[*Conn]bool),
	}
}

var h = NewHub()

func Run() {
	h.Run()
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.register:
			log.Println("register")
			h.connections[conn] = true
		case conn := <-h.unregister:
			log.Println("unregister")
			if _, ok := h.connections[conn]; ok {
				delete(h.connections, conn)
				close(conn.Send)
			}
		case message := <-h.broadcast:
			log.Println("broadcast")
			for conn := range h.connections {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(h.connections, conn)
				}
			}
		}
	}
}

func (h *Hub) Register(conn *Conn) {
	h.register <- conn
}

func (h *Hub) UnRegister(conn *Conn) {
	h.unregister <- conn
}

func (h *Hub) Broadcast(msg *Message) {
	h.broadcast <- msg
}
