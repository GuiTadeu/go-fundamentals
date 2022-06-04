package exercises

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id        uint64
	Code      string
	Currency  string
	Value     float64
	Issuer    string
	Receiver  string
	CreatedAt string
}

func (t *Transaction) match(params url.Values) (match bool) {

	if (params.Get("id") != "") && (fmt.Sprint(t.Id) == params.Get("id")) {
		match = true

		if len(params) == 1 {
			return true
		}
	} else if params.Get("id") != "" && len(params) > 1 {
		return false
	}

	if (params.Get("code") != "") && (t.Code == params.Get("code")) {
		match = true

		if len(params) == 1 {
			return true
		}
	} else if params.Get("code") != "" && len(params) > 1 {
		return false
	}

	if (params.Get("currency") != "") && (t.Currency == params.Get("currency")) {
		match = true

		if len(params) == 1 {
			return true
		}
	} else if params.Get("currency") != "" && len(params) > 1 {
		return false
	}

	if (params.Get("value") != "") && (fmt.Sprint(t.Value) == params.Get("value")) {
		match = true

		if len(params) == 1 {
			return true
		}
	} else if params.Get("value") != "" && len(params) > 1 {
		return false
	}

	if (params.Get("issuer") != "") && (t.Issuer == params.Get("issuer")) {
		match = true

		if len(params) == 1 {
			return true
		}
	} else if params.Get("issuer") != "" && len(params) > 1 {
		return false
	}

	if (params.Get("receiver") != "") && (t.Receiver == params.Get("receiver")) {
		match = true

		if len(params) == 1 {
			return true
		}
	} else if params.Get("receiver") != "" && len(params) > 1 {
		return false
	}

	if (params.Get("createdAt") != "") && (t.CreatedAt == params.Get("createdAt")) {
		match = true

		if len(params) == 1 {
			return true
		}
	} else if params.Get("createdAt") != "" && len(params) > 1 {
		return false
	}

	return
}

var transactions = []Transaction{
	{
		Id:        1,
		Code:      "EUAPAY",
		Currency:  "USD",
		Value:     42.00,
		Issuer:    "pica-pau",
		Receiver:  "pé-de-pano",
		CreatedAt: "2022-06-02",
	},
	{
		Id:        2,
		Code:      "BRPAY",
		Currency:  "BRL",
		Value:     66.00,
		Issuer:    "pé-de-pano",
		Receiver:  "pica-pau",
		CreatedAt: "2022-06-02",
	},
	{
		Id:        3,
		Code:      "EUAPAY",
		Currency:  "USD",
		Value:     120.00,
		Issuer:    "pica-pau",
		Receiver:  "pé-de-pano",
		CreatedAt: "2022-06-02",
	},
}

func FilterHandler(ctx *gin.Context) {

	params := ctx.Request.URL.Query()
	result := []Transaction{}

	for _, t := range transactions {
		if t.match(params) {
			result = append(result, t)
		}
	}

	ctx.JSON(200, result)
	return
}

func GetOneHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	for _, transaction := range transactions {
		if fmt.Sprint(transaction.Id) == id {
			ctx.JSON(200, transaction)
			return
		}
	}

	ctx.JSON(404, "not found transaction")
	return
}
