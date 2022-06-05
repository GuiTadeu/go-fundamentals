package handler

import (
	"github.com/GuiTadeu/meli-go/web/class06/internal/transactions"
	"github.com/GuiTadeu/meli-go/web/class06/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateTransactionRequest struct {
	Code      string  `json:"code" binding:"required"`
	Currency  string  `json:"currency" binding:"required"`
	Value     float64 `json:"value" binding:"required"`
	Issuer    string  `json:"issuer" binding:"required"`
	Receiver  string  `json:"receiver" binding:"required"`
	CreatedAt string  `json:"createdAt" binding:"required"`
}

type TransactionHandler struct {
	serv transactions.TransactionService
}

func NewTransactionHandler(serv transactions.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		serv: serv,
	}
}

// GetAll godoc
// @Summary List Transactions
// @Description Get All Transactions
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (handler *TransactionHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "GOIABA" {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}

		transactions, err := handler.serv.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, "internal server error"))
			return
		}

		ctx.JSON(200, web.NewResponse(200, transactions, ""))
	}
}

// Store godoc
// @Summary Save Transaction
// @Description Create Transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /transactions [post]
func (handler *TransactionHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "GOIABA" {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}

		var request CreateTransactionRequest
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "field error"))
			return
		}

		transaction, err := handler.serv.Store(
			request.Code, request.Currency, request.Value,
			request.Issuer, request.Receiver, request.CreatedAt,
		)

		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, "failed to save"))
			return
		}

		ctx.JSON(201, web.NewResponse(201, transaction, ""))
	}
}