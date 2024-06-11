package routes

import (
	"WeMeet/api/controllers"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(r *gin.Engine) {
	r.GET("/room/:id", controllers.RoomController)
}
