package exercises

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var ID uint64 = 0

const TOKEN = "GOIABA"

var transactions = []Transaction{}

// Ex01 - POST
func CreateTransactionHandler(ctx *gin.Context) {

	// Ex03 - HEADER
	token := ctx.GetHeader("token")
	if token != TOKEN {
		ctx.JSON(401, gin.H{
			"error": "not authorized",
		})
		return
	}

	var transaction Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ID++
	transaction.Id = ID

	transactions = append(transactions, transaction)
	ctx.JSON(201, transaction)
	fmt.Println(transactions)
}

// Ex02 - Required
type Transaction struct {
	Id        uint64  `json:"id"`
	Code      string  `json:"code" binding:"required"`
	Currency  string  `json:"currency" binding:"required"`
	Value     float64 `json:"value" binding:"required"`
	Issuer    string  `json:"issuer" binding:"required"`
	Receiver  string  `json:"receiver" binding:"required"`
	CreatedAt string  `json:"createdAt" binding:"required"`
}
