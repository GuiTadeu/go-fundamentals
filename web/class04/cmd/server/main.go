package main

import (
	"github.com/GuiTadeu/meli-go/web/class04/cmd/server/handler"
	"github.com/GuiTadeu/meli-go/web/class04/internal/transactions"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := transactions.NewTransactionRepo()
	service := transactions.NewTransactionServ(repo)
	handler := handler.NewTransactionHandler(service)

	server := gin.Default()

	transactionsGroup := server.Group("/transactions")
	transactionsGroup.POST("/", handler.Store())
	transactionsGroup.GET("/", handler.GetAll())

	server.Run()
}