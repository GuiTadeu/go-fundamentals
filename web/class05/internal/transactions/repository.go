package transactions

import "github.com/GuiTadeu/meli-go/web/class05/pkg/store"

var ID uint64

type Transaction struct {
	Id        uint64
	Code      string
	Currency  string
	Value     float64
	Issuer    string
	Receiver  string
	CreatedAt string
}

type TransactionRepository interface {
	Store(code string, currency string, value float64, issuer string, receiver string, createdAt string) (Transaction, error)
	GetAll() ([]Transaction, error)
}

type TransactionRepositoryImpl struct {
	db store.Store
}

func NewTransactionRepository(db store.Store) TransactionRepository {
	return &TransactionRepositoryImpl {
		db: db,
	}
}

func (repo *TransactionRepositoryImpl) Store(
	code string, currency string, value float64, 
	issuer string, receiver string, createdAt string,
) (Transaction, error) {
	var transactions []Transaction
	repo.db.Read(&transactions)

	ID++
	transaction := Transaction{ID, code, currency, value, issuer, receiver, createdAt}
	transactions = append(transactions, transaction)

	if err := repo.db.Write(transactions); err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func (repo *TransactionRepositoryImpl) GetAll() ([]Transaction, error) {
	var transactions []Transaction
	repo.db.Read(&transactions)
	return transactions, nil
}