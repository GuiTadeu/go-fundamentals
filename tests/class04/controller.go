package class04

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transactionService TransactionService
}

func NewTransactionController(s TransactionService) *transactionController {
	return &transactionController{
		transactionService: s,
	}
}

func (c *transactionController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request Transaction

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(
				http.StatusUnprocessableEntity,
				NewResponse(http.StatusUnprocessableEntity, nil, err.Error()),
			)
			return
		}

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				NewResponse(http.StatusBadRequest, nil, "transaction id binding error"),
			)
			return
		}

		updatedTransaction, err := c.transactionService.Update(
			id,
			request.Code,
			request.Currency,
			request.Value,
			request.Issuer,
			request.Receiver,
			request.CreatedAt,
		)

		if err != nil {
			ctx.JSON(500, NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, NewResponse(http.StatusOK, updatedTransaction, ""))
	}
}

func (c *transactionController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(
				http.StatusNotFound,
				NewResponse(http.StatusNotFound, nil, err.Error()),
			)
			return
		}

		err = c.transactionService.Delete(id)

		if err != nil {
			ctx.JSON(500, NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusNoContent, NewResponse(http.StatusNoContent, nil, ""))
	}
}