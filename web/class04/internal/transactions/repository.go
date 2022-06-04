package transactions

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
