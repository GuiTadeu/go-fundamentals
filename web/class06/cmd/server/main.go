package main

import (
	"github.com/GuiTadeu/meli-go/web/class06/cmd/server/handler"
	"github.com/GuiTadeu/meli-go/web/class06/internal/transactions"
	"github.com/gin-gonic/gin"
	"github.com/GuiTadeu/meli-go/web/class06/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Transactions
func main() {

	repo := transactions.NewTransactionRepo()
	service := transactions.NewTransactionServ(repo)
	handler := handler.NewTransactionHandler(service)

	server := gin.Default()

	docs.SwaggerInfo.Host = "localhost:8080"
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	transactionsGroup := server.Group("/transactions")
	transactionsGroup.POST("/", handler.Store())
	transactionsGroup.GET("/", handler.GetAll())

	server.Run()
}