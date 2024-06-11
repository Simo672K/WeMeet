package chatservice

import "golang.org/x/net/websocket"



type ChatClient struct {
	SocketConn *websocket.Conn
	Send   chan []byte
	Room       *ChatRoom
}

func (c *ChatClient) WriteSocket() {
	defer c.SocketConn.Close()
}