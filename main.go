package main

import (
	"github.com/gin-gonic/gin"
	"untitled/handler"
)

func main() {
	router := gin.Default()

	router.GET("/items", handler.GetItem)
	router.GET("/members", handler.GetMember)

	router.Run(":8080")
}
