package hub

type Hub struct {
	// Registered connections.
	clients map[*Client]bool

	// Inbound messages from the connections.
	broadcast chan *Message

	// Register requests from the connections.
	register chan *Client

	// Unregister requests from connections.
	unregister chan *Client

	MessageReceived func(receivedMessage *Message, receivedClient *Client)
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
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

func (h *Hub) SendToUser(userID string, msg *Message) {
	for client := range h.clients {
		if client.UserID == userID {
			client.Send <- msg
		}
	}
}
