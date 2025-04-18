package websocket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		client := make(ClientChan, 256)
		hub.Register(client)
		defer hub.UnRegister(client)

		go func() {
			fmt.Printf("Working")
			for msg := range client {
				fmt.Printf("📤 Sending to client: %s\n", msg)

				err := conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					fmt.Printf("❌ Error writing message: %v\n", err)
					return
				}
			}
		}()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("❌ Error reading message:", err)
				break
			}
			fmt.Printf("📥 Received from client: [%d] %s\n", messageType, message)
		}
	}
}
