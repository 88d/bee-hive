package hub

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var ids = 0

func (h *Hub) ServeHub() http.HandlerFunc {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		ids += 1
		conn := NewClient(ws, string(ids), "blablabla")
		h.register <- conn
		go conn.writePump()
		conn.readPump()
	}
}

func ServeHub() http.HandlerFunc {
	return h.ServeHub()
}
