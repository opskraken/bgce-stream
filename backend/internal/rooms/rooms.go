package rooms

import "sync"

type Room struct {
	ID      string
	Members map[string]bool
}

type RoomManager struct {
	rooms map[string]*Room
	mu    sync.Mutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*Room),
	}
}

func (rm *RoomManager) CreateRoom(id string) *Room {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	room := &Room{
		ID:      id,
		Members: make(map[string]bool),
	}
	rm.rooms[id] = room
	return room
}

func (rm *RoomManager) DeleteRoom(id string) (*Room, bool) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	r, ok := rm.rooms[id]
	return r, ok
}
