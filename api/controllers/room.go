package controllers

import (
	"WeMeet/config"
	"WeMeet/internals/repository"

	"github.com/gin-gonic/gin"
)

func RoomController(c *gin.Context) {
	dbInstance := config.DBConnect()
	roomRepo := repository.RoomRepository{
		DBClient: dbInstance,
	}

	roomRepo.CreateRoom()
	c.String(200, "Running redis test if it works")
}
