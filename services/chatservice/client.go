package chatservice

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type ChatClient struct {
	SocketConn *websocket.Conn
	Send       chan []byte
	Room       *ChatRoom
}

func (c *ChatClient) ReadMessage() {
	defer c.SocketConn.Close()

	fmt.Println("Client go routine is running...")
	for {
		_, msg, err := c.SocketConn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseMessage) {
				fmt.Println("Client disconnected")
			} else {
				log.Fatal("***",err)
			}
			break
		}
		fmt.Println(string(msg))
		c.Room.Forward <- msg
	}
}
