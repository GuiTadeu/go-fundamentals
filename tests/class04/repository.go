package class04

import "errors"

var (
	TransactionNotFoundError   = errors.New("transaction not found")
)

var transactions []Transaction
var ID uint64

type TransactionRepository interface {
	Update(id uint64, code string, currency string, value float64, issuer string, receiver string, createdAt string) (Transaction, error)
	Delete(id uint64) error
}

type TransactionRepositoryImpl struct{}

func NewTransactionRepo() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (r *TransactionRepositoryImpl) Update(
	id uint64, code string, currency string, value float64,
	issuer string, receiver string, createdAt string,
) (Transaction, error) {

	transaction := Transaction{
		Code:      code,
		Currency:  currency,
		Value:     value,
		Issuer:    issuer,
		Receiver:  receiver,
		CreatedAt: createdAt,
	}

	for i := range transactions {
		if transactions[i].Id == id {
			transaction.Id = id
			transactions[i] = transaction
		} else {
			return Transaction{}, TransactionNotFoundError
		}
	}

	return transaction, nil
}

func (r *TransactionRepositoryImpl) Delete(id uint64) error {

	var deleted bool
	var index int
	for i := range transactions {
		if transactions[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return TransactionNotFoundError
	}

	transactions = append(transactions[:index], transactions[index+1:]...)
	return nil
}
