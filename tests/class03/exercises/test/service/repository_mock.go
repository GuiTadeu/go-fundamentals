package exercises

import (
	"fmt"

	itrl "github.com/GuiTadeu/meli-go/tests/class03/exercises/internal"
)

var transactions = []itrl.Transaction{
	{Id:1, Code:"oldCode", Currency:"oldCurrency", Value:1.0, Issuer:"oldIssuer", Receiver:"oldReceiver", CreatedAt:"2012-12-12"},
	{Id:2, Code:"oldCode", Currency:"oldCurrency", Value:1.0, Issuer:"oldIssuer", Receiver:"oldReceiver", CreatedAt:"2012-12-12"},
}
var ID uint64

type MockRepository struct {
	UpdateWasCalled bool
	GetAllWasCalled bool
}

func (r *MockRepository) GetAll() ([]itrl.Transaction, error) {
	r.GetAllWasCalled = true
	return transactions, nil
}

func (m *MockRepository) Update(
	id uint64, code string, currency string, value float64,
	issuer string, receiver string, createdAt string,
) (itrl.Transaction, error) {

	m.UpdateWasCalled = true

	transaction := itrl.Transaction{
		Code:      code,
		Currency:  currency,
		Value:     value,
		Issuer:    issuer,
		Receiver:  receiver,
		CreatedAt: createdAt,
	}

	allTransactions, _ := m.GetAll()

	for i, t := range allTransactions {
		if t.Id == id {
			transaction.Id = id
			transactions[i] = transaction
			return transaction, nil
		}
	}

	return itrl.Transaction{}, fmt.Errorf("not found transaction")
}

func (r *MockRepository) Delete(id uint64) error {

	var deleted bool
	var index int

	allTransactions, _ := r.GetAll()

	for i := range allTransactions {
		if transactions[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("not found transaction")
	}

	transactions = append(transactions[:index], transactions[index+1:]...)
	return nil
}
