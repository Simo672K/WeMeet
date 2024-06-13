package chatservice

import (
	"fmt"
	"log"
	"net/http"

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
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (cr *ChatRoom) StartChatService(w http.ResponseWriter, r *http.Request) {
	var conn *websocket.Conn
	var err error
	conn = cr.RoomWSConn

	if(cr.RoomWSConn == nil ){
		fmt.Println("Init upgrader")
		conn, err = Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal("Failed to setting up a websockets server")
		}
		cr.RoomWSConn = conn
	}

	defer conn.Close()

	client := ChatClient{
		SocketConn: conn,
		Send:       make(chan []byte, MsgBufferSize),
		Room:       cr,
	}
	go client.ReadMessage()

	go cr.RunRoom()
	cr.Join <- &client
}
