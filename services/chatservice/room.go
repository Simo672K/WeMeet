package chatservice

import "time"

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
}

func (cr *ChatRoom) InitRoomChannels() {
	rsa := make(chan bool) // room still active channel
	frwrd := make(chan []byte)
	join := make(chan *ChatClient)
	leave := make(chan *ChatClient)
	actvClients := make(map[*ChatClient]bool)

	cr.RoomStillActive = rsa
	cr.Forward = frwrd
	cr.Join = join
	cr.Leave = leave
	cr.ActiveClients = actvClients
}

func (cr *ChatRoom) RunRoom() {
	for {
		select {
			case client := <- cr.Join:
				cr.ActiveClients[client]= true
			case client := <- cr.Leave:
				delete(cr.ActiveClients, client)
			case msg := <- cr.Forward:
				for client := range cr.ActiveClients {
					client.Send <- msg
				} 

		}
	}
}
