package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
		panic("An error accured during app execution, pannicking...")
	}
}
