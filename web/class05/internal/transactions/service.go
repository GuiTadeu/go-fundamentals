package transactions

type TransactionService interface {
	GetAll() ([]Transaction, error)
	Store(code string, currency string, value float64, issuer string, receiver string, createdAt string) (Transaction, error)
}

type TransactionServiceImpl struct {
	repo TransactionRepository
}

func NewTransactionServ(repo TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		repo: repo,
	}
}

func (serv *TransactionServiceImpl) GetAll() ([]Transaction, error) {
	transactions, err := serv.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (serv *TransactionServiceImpl) Store(
	code string, currency string, value float64, 
	issuer string, receiver string, createdAt string,
) (Transaction, error) {
	
	transaction, err := serv.repo.Store(
		code, currency, value, 
		issuer, receiver, createdAt,
	)
	
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}