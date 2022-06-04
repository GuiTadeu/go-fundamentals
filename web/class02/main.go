package main

import (
	"github.com/GuiTadeu/meli-go/web/class02/exercises"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/transactions", exercises.FilterHandler)
	server.GET("/transactions/:id", exercises.GetOneHandler)
	server.Run(":8080")
}