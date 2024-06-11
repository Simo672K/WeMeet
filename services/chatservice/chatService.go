package chatservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ActiveRooms struct {
	Rooms []ChatRoom
}

type Message struct {
	Content string `json:"content"`
}

const (
	BufferSize    = 1024
	MsgBufferSize = 256
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  BufferSize,
	WriteBufferSize: MsgBufferSize,
	CheckOrigin: func(r *http.Request) bool { return true },
}


func RunChatService(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Failed to setting up a websockets server")
	}
	defer conn.Close()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)

		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				fmt.Println("Connection closed by client")
				break
			} else {
				fmt.Println("Error reading message:", err)
				break
			}
		}
		fmt.Printf("Received message: %+v\n", msg)

		if err := conn.WriteJSON(gin.H{
			"message": "Welcome",
		}); err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

	}
}
