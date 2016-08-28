package hub

import (
	"log"
	"net/http"
	"strconv"

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
		ids++
		client := NewClient(ws, strconv.Itoa(ids), "blablabla")
		h.register <- client
		go client.writePump()
		client.readPump()
	}
}

func ServeHub() http.HandlerFunc {
	return root.ServeHub()
}
