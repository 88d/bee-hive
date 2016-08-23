package hub

type Hub struct {
	// Registered connections.
	connections map[*Client]bool

	// Inbound messages from the connections.
	broadcast chan *Message

	// Register requests from the connections.
	register chan *Client

	// Unregister requests from connections.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:   make(chan *Message),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		connections: make(map[*Client]bool),
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
			h.connections[conn] = true
		case conn := <-h.unregister:
			if _, ok := h.connections[conn]; ok {
				delete(h.connections, conn)
				close(conn.Send)
			}
		case message := <-h.broadcast:
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

func (h *Hub) Register(conn *Client) {
	h.register <- conn
}

func (h *Hub) UnRegister(conn *Client) {
	h.unregister <- conn
}

func (h *Hub) Broadcast(msg *Message) {
	h.broadcast <- msg
}
