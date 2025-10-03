package signaling

import (
	"log"
	"net/http"

	"bgce-stream/internal/rooms"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request, rm *rooms.RoomManager) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	log.Println("New WebSocket connection established")

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		log.Printf("Received message: %s\n", string(msg))

		// Echo back for now
		if err := conn.WriteMessage(mt, msg); err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
