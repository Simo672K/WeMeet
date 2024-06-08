package bootstrap

import (
	"WeMeet/api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	routes.ChatRoutes(router)
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
		panic("An error accured during app execution, pannicking...")
	}
}
