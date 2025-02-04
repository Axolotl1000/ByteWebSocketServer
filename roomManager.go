package main

import (
	"log"

	"github.com/gorilla/websocket"
)

var (
	connections = make(map[string]map[string]*websocket.Conn)
)

func addConnection(roomID, userID string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := connections[roomID]; !exists {
		connections[roomID] = make(map[string]*websocket.Conn)
	}
	connections[roomID][userID] = conn
}

func removeConnection(roomID, userID string) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := connections[roomID]; exists {
		delete(connections[roomID], userID)

		if len(connections[roomID]) == 0 {
			delete(connections, roomID)
		}
	}
}

func broadcastToRoom(roomID, senderID string, msg []byte) {
	mu.Lock()
	defer mu.Unlock()

	if room, exists := connections[roomID]; exists {
		for userID, conn := range room {
			if userID != senderID {
				err := conn.WriteMessage(websocket.BinaryMessage, msg)
				if err != nil {
					log.Println("廣播錯誤:", err)
				}
			}
		}
	}
}
