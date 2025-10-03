package main

import (
	"bgce-stream/internal/rooms"
	"bgce-stream/internal/signaling"
	"log"
	"net/http"
)

func main() {
	rm := rooms.NewRoomManager()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		signaling.HandleWebSocket(w, r, rm)
	})

	log.Println("bgce-stream backend running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
