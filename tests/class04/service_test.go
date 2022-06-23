package class04

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {

	expectedResult := Transaction{
		Id: 1, Code: "newCode", Currency: "newCurrency", Value: 1.0,
		Issuer: "newIssuer", Receiver: "newReceiver", CreatedAt: "2012-12-12",
	}

	mockRepository := mockTransactionRepository{
		result: expectedResult,
		err:    nil,
	}

	service := NewTransactionServ(mockRepository)
	result, err := service.Update(1, "newCode", "newCurrency", 1.0, "newIssuer", "newReceiver", "2012-12-12")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.True(t, UpdateWasCalledOnRepository)
}

func TestDelete(t *testing.T) {

	mockRepository := mockTransactionRepository{
		result: Transaction{},
		err:    nil,
	}

	service := NewTransactionServ(mockRepository)
	err := service.Delete(1)

	assert.Nil(t, err)
	assert.True(t, DeleteWasCalledOnRepository)
}

func TestDeleteError(t *testing.T) {

	expectedError := errors.New("KKK")
	mockRepository := mockTransactionRepository{
		result: Transaction{},
		err:    expectedError,
	}

	service := NewTransactionServ(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, expectedError, err)
	assert.True(t, DeleteWasCalledOnRepository)
}
