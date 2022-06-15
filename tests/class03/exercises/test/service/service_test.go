package exercises

import (
	"testing"

	itrl "github.com/GuiTadeu/meli-go/tests/class03/exercises/internal"
	"github.com/stretchr/testify/assert"
)

// Not is FIRST
func TestUpdate(t *testing.T) {
	mockService := MockService{}
	expectedResult := itrl.Transaction{
		Id: 1, Code: "newCode", Currency: "newCurrency", Value: 1.0,
		Issuer: "newIssuer", Receiver: "newReceiver", CreatedAt: "2012-12-12",
	}

	result, _ := mockService.Update(1, "newCode", "newCurrency", 1.0, "newIssuer", "newReceiver", "2012-12-12")

	assert.Equal(t, expectedResult, result)
	assert.True(t, mockService.repo.GetAllWasCalled)
	assert.True(t, mockService.repo.UpdateWasCalled)
}

// Not is FIRST
func TestDelete(t *testing.T) {
	mockService := MockService{}

	transactions, _ := mockService.repo.GetAll()
	assert.Len(t, transactions, 2)

	_ = mockService.Delete(1)
	transactions, _ = mockService.repo.GetAll()

	assert.Len(t, transactions, 1)
	assert.True(t, mockService.repo.GetAllWasCalled)
}

// Not is FIRST
func TestDeleteError(t *testing.T) {
	mockService := MockService{}

	expectedError := "not found transaction"
	error := mockService.Delete(100)

	assert.Equal(t, expectedError, error.Error())
}