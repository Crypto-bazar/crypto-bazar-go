package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

type ClientChan chan []byte

type Hub struct {
	mu      sync.Mutex
	clients map[ClientChan]bool
}

func NewHub() *Hub {
	return &Hub{clients: make(map[ClientChan]bool)}
}

func (h *Hub) Register(client ClientChan) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[client] = true
}

func (h *Hub) UnRegister(client ClientChan) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.clients[client]; ok {
		delete(h.clients, client)
		close(client)
	}
}

func (h *Hub) Broadcast(v any) {
	h.mu.Lock()
	defer h.mu.Unlock()

	data, err := json.Marshal(v)
	if err != nil {
		log.Println("Broadcast marshal error:", err)
		return
	}

	fmt.Printf("Broadcast message: %s\n", data)

	for ch := range h.clients {
		select {
		case ch <- data:
			log.Println("Message sent to a client")
		default:
			log.Println("Client channel full or not reading, removing")
			delete(h.clients, ch)
			close(ch)
		}
	}
}
