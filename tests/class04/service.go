package class04

type TransactionService interface {
	Update(id uint64, code string, currency string, value float64, issuer string, receiver string, createdAt string) (Transaction, error)
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

func (serv *TransactionServiceImpl) Update(
	id uint64, code string, currency string, value float64,
	issuer string, receiver string, createdAt string,
) (Transaction, error) {
	return serv.repo.Update(id, code, currency, value, issuer, receiver, createdAt)
}

func (serv *TransactionServiceImpl) Delete(id uint64) error {
	return serv.repo.Delete(id)
}
