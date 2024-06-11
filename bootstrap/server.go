package bootstrap

import (
	"WeMeet/api/routes"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()
	router.Use(cors.Default())

	routes.ChatRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
		panic("An error accured during app execution, pannicking...")
	}
}
