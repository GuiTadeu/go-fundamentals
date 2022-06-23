package main

import (
	"github.com/GuiTadeu/meli-go/tests/class04"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	transactionRepository := class04.NewTransactionRepo()
	transactionService := class04.NewTransactionServ(transactionRepository)
	transactionController := class04.NewTransactionController(transactionService)

	transactionGroup := server.Group("/api/v1/transactions")
	transactionGroup.PATCH("/:id", transactionController.Update())
	transactionGroup.DELETE("/:id", transactionController.Delete())

	server.Run()
}
