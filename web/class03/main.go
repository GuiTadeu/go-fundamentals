package main

import (
	"github.com/GuiTadeu/meli-go/web/class03/exercises"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.POST("/transactions", exercises.CreateTransactionHandler)
	server.Run(":8080")
}