package exercises

type TransactionService interface {
	GetAll() ([]Transaction, error)
	Store(code string, currency string, value float64, issuer string, receiver string, createdAt string) (Transaction, error)
	Update(id uint64, code string, currency string, value float64, issuer string, receiver string, createdAt string) (Transaction, error)
	UpdateCode(id uint64, code string) (Transaction, error)
	Delete(id uint64) error
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

func (serv *TransactionServiceImpl) Update(
	id uint64, code string, currency string, value float64,
	issuer string, receiver string, createdAt string,
) (Transaction, error) {
	return serv.repo.Update(id, code, currency, value, issuer, receiver, createdAt)
}

func (serv *TransactionServiceImpl) UpdateCode(id uint64, code string) (Transaction, error) {
	return serv.repo.UpdateCode(id, code)
}

func (serv *TransactionServiceImpl) Delete(id uint64) error {
	return serv.repo.Delete(id)
}
