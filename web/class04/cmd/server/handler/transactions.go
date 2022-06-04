package handler

import (
	"github.com/GuiTadeu/meli-go/web/class04/internal/transactions"
	"github.com/gin-gonic/gin"
)

type CreateTransactionRequest struct {
	Id        uint64  `json:"id"`
	Code      string  `json:"code" binding:"required"`
	Currency  string  `json:"currency" binding:"required"`
	Value     float64 `json:"value" binding:"required"`
	Issuer    string  `json:"issuer" binding:"required"`
	Receiver  string  `json:"receiver" binding:"required"`
	CreatedAt string  `json:"createdAt" binding:"required"`
}

type TransactionHandlerImpl struct {
	serv transactions.TransactionService
}

func NewTransactionHandler(serv transactions.TransactionService) *TransactionHandlerImpl {
	return &TransactionHandlerImpl{
		serv: serv,
	}
}

func (handler *TransactionHandlerImpl) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "GOIABA" {
			ctx.JSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		transactions, err := handler.serv.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, transactions)
	}
}

func (handler *TransactionHandlerImpl) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "GOIABA" {
			ctx.JSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		var request CreateTransactionRequest
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(404, gin.H {
				"error": err.Error(),
			})
			return
		}

		transaction, err := handler.serv.Store(
			request.Code, request.Currency, request.Value, 
			request.Issuer, request.Receiver, request.CreatedAt,
		)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(201, transaction)
	}
}