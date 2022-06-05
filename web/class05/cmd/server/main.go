package main

import (
	"github.com/GuiTadeu/meli-go/web/class05/cmd/server/handler"
	"github.com/GuiTadeu/meli-go/web/class05/internal/transactions"
	"github.com/GuiTadeu/meli-go/web/class05/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	db := store.New(store.FileType, "./transactions.json")
	repo := transactions.NewTransactionRepository(db)
	service := transactions.NewTransactionServ(repo)
	handler := handler.NewTransactionHandler(service)

	server := gin.Default()

	transactionsGroup := server.Group("/transactions")
	transactionsGroup.POST("/", handler.Store())
	transactionsGroup.GET("/", handler.GetAll())

	server.Run()
}