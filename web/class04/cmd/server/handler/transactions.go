package handler

import (
	"strconv"

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

type UpdateTransactionRequest struct {
	Code      string  `json:"code" binding:"required"`
	Currency  string  `json:"currency" binding:"required"`
	Value     float64 `json:"value" binding:"required"`
	Issuer    string  `json:"issuer" binding:"required"`
	Receiver  string  `json:"receiver" binding:"required"`
	CreatedAt string  `json:"createdAt" binding:"required"`
}

type UpdateTransactionCodeRequest struct {
	Code string `json:"code" binding:"required"`
}

type TransactionHandler struct {
	serv transactions.TransactionService
}

func NewTransactionHandler(serv transactions.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		serv: serv,
	}
}

func (handler *TransactionHandler) GetAll() gin.HandlerFunc {
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

func (handler *TransactionHandler) Store() gin.HandlerFunc {
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
			ctx.JSON(400, gin.H{
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

func (handler *TransactionHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "GOIABA" {
			ctx.JSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid ID",
			})
		}

		var request UpdateTransactionRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		transaction, err := handler.serv.Update(id, request.Code, request.Currency,
			request.Value, request.Issuer, request.Receiver, request.CreatedAt)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, transaction)
	}
}

func (handler *TransactionHandler) UpdateCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "GOIABA" {
			ctx.JSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid ID",
			})
		}

		var request UpdateTransactionCodeRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		transaction, err := handler.serv.UpdateCode(id, request.Code)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, transaction)
	}
}

func (handler *TransactionHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "GOIABA" {
			ctx.JSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid ID",
			})
			return
		}

		err = handler.serv.Delete(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{"data": "transaction was successfully removed!"})
	}
}
