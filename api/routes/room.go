package routes

import (
	// "WeMeet/api/controllers"
	"WeMeet/services/chatservice"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(r *gin.Engine) {
	room := chatservice.ChatRoom{}
	r.GET("/room/:id", room.RoomController)
}
