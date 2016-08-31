package hub

import (
	"log"
	"net/http"

	"github.com/black-banana/bee-hive/hive/auth"
	"github.com/gorilla/websocket"
)

func (h *Hub) ServeHub() http.HandlerFunc {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := auth.ClaimsFromRequestQuery(r)
		if err != nil {
			log.Print(err)
			w.WriteHeader(400)
			return
		}

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := NewClient(ws, claims.UserID, h)
		h.register <- client
		go client.writePump()
		client.readPump()
	}
}
