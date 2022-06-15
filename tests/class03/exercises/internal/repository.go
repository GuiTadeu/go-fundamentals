package exercises

import "fmt"

type Transaction struct {
	Id        uint64
	Code      string
	Currency  string
	Value     float64
	Issuer    string
	Receiver  string
	CreatedAt string
}

var transactions []Transaction
var ID uint64

type TransactionRepository interface {
	GetAll() ([]Transaction, error)
	Store(code string, currency string, value float64, issuer string, receiver string, createdAt string) (Transaction, error)
	Update(id uint64, code string, currency string, value float64, issuer string, receiver string, createdAt string) (Transaction, error)
	UpdateCode(id uint64, code string) (Transaction, error)
	Delete(id uint64) error
}

type TransactionRepositoryImpl struct{}

func NewTransactionRepo() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (r *TransactionRepositoryImpl) GetAll() ([]Transaction, error) {
	return transactions, nil
}

func (r *TransactionRepositoryImpl) Store(
	code string, currency string, value float64,
	issuer string, receiver string, createdAt string,
) (Transaction, error) {
	ID++
	newTransaction := Transaction{ID, code, currency, value, issuer, receiver, createdAt}
	transactions = append(transactions, newTransaction)
	return newTransaction, nil
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
			return Transaction{}, fmt.Errorf("not found transaction")
		}
	}

	return transaction, nil
}

func (r *TransactionRepositoryImpl) UpdateCode(id uint64, code string) (Transaction, error) {

	var transaction Transaction
	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Code = code
			transaction = transactions[i]
		} else {
			return Transaction{}, fmt.Errorf("not found transaction")
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
		return fmt.Errorf("not found transaction")
	}

	transactions = append(transactions[:index], transactions[index+1:]...)
	return nil
}
