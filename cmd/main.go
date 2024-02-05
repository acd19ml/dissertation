package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//getting context
	defer Disconnect()
	router := gin.Default()
	router.POST("/books", CreateBook)
	router.GET("/books", ListBooks)
	router.GET("/books/:name", FindBook)
	router.Run("localhost:8080")
}
