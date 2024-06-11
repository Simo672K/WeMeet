package controllers

import (
	// "WeMeet/config"
	// "WeMeet/internals/repository"

	"WeMeet/services/chatservice"

	"github.com/gin-gonic/gin"
)

func RoomController(c *gin.Context) {
	// dbInstance := config.DBConnect()
	// roomRepo := repository.RoomRepository{
	// 	DBClient: dbInstance,
	// }

	// roomRepo.CreateRoom("test")
	chatservice.RunChatService(c.Writer, c.Request)
	// c.String(200, "Room service")
}
