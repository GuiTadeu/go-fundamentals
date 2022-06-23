package class04

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Update_200(t *testing.T) {

	transactionToUpdate := Transaction{
		Id:        1,
		Code:      "ABC",
		Currency:  "ABC",
		Value:     1,
		Issuer:    "ABC",
		Receiver:  "ABC",
		CreatedAt: "ABC",
	}

	jsonValue, _ := json.Marshal(transactionToUpdate)
	requestBody := bytes.NewBuffer(jsonValue)

	updatedTransaction := Transaction{
		Id:        1,
		Code:      "DEF",
		Currency:  "DEF",
		Value:     1,
		Issuer:    "DEF",
		Receiver:  "DEF",
		CreatedAt: "DEF",
	}

	mockService := mockTransactionService{
		result: updatedTransaction,
		err:    nil,
	}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/v1/transactions/1", requestBody)
	router.ServeHTTP(response, request)

	responseData := Transaction{}
	decodeWebResponse(response, &responseData)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, updatedTransaction, responseData)
}

func decodeWebResponse(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func Test_Delete_204(t *testing.T) {

	mockService := mockTransactionService{
		result: Transaction{},
		err:    nil,
	}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/transactions/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 204, response.Code)
}

func Test_Delete_404(t *testing.T) {

	mockService := mockTransactionService{
		result: Transaction{},
		err:    TransactionNotFoundError,
	}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/products/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)
}

func setupRouter(mockService mockTransactionService) *gin.Engine {
	controller := NewTransactionController(mockService)

	router := gin.Default()

	router.PATCH("/api/v1/transactions/:id", controller.Update())
	router.DELETE("/api/v1/transactions/:id", controller.Delete())

	return router
}
