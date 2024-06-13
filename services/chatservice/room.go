package chatservice

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatRoomMetaData struct {
	Id         string
	RoomName   string
	ExpiresIn  int
	MaxClients int
	CreatedAt  *time.Time
}

type ChatRoom struct {
	ChatRoomMetaData
	RoomStillActive chan bool
	Forward         chan []byte
	Join            chan *ChatClient
	Leave           chan *ChatClient
	ActiveClients   map[*ChatClient]bool
	RoomWSConn     *websocket.Conn
}

func (cr *ChatRoom) InitRoomChannels() {
	rsa := make(chan bool, 1) // room still active channel
	frwrd := make(chan []byte)
	join := make(chan *ChatClient)
	leave := make(chan *ChatClient)
	actvClients := make(map[*ChatClient]bool)

	rsa <- true
	cr.RoomStillActive = rsa
	cr.Forward = frwrd
	cr.Join = join
	cr.Leave = leave
	cr.ActiveClients = actvClients
}

func (cr *ChatRoom) RunRoom() {
	cr.InitRoomChannels()

	fmt.Println("Chat Room go routine is running...")
	for {
		select {
		case client := <-cr.Join:
			cr.ActiveClients[client] = true
		case client := <-cr.Leave:
			delete(cr.ActiveClients, client)
		case msg := <-cr.Forward:
			for client := range cr.ActiveClients {
				client.Send <- msg
			}

		}
	}
}

func (cr *ChatRoom) RoomController(c *gin.Context) {
	// dbInstance := config.DBConnect()
	// roomRepo := repository.RoomRepository{
	// 	DBClient: dbInstance,
	// }

	// roomRepo.CreateRoom("test")
	cr.StartChatService(c.Writer, c.Request)
	// c.String(200, "Room service")
}
