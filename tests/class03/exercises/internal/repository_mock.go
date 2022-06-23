package internal

var DeleteWasCalledOnRepository bool
var UpdateWasCalledOnRepository bool

type mockTransactionRepository struct {
	result any
	err    error
}

func (m mockTransactionRepository) GetAll() ([]Transaction, error) {

	if m.err != nil {
		return []Transaction{}, m.err
	}
	return m.result.([]Transaction), nil
}

func (m mockTransactionRepository) Store(
	code string, currency string, value float64,
	issuer string, receiver string, createdAt string,
) (Transaction, error) {

	if m.err != nil {
		return Transaction{}, m.err
	}
	return m.result.(Transaction), nil
}

func (m mockTransactionRepository) Update(
	id uint64, code string, currency string, value float64,
	issuer string, receiver string, createdAt string,
) (Transaction, error) {

	UpdateWasCalledOnRepository = true

	if m.err != nil {
		return Transaction{}, m.err
	}
	return m.result.(Transaction), nil
}

func (m mockTransactionRepository) UpdateCode(id uint64, code string) (Transaction, error) {
	if m.err != nil {
		return Transaction{}, m.err
	}
	return m.result.(Transaction), nil
}

func (m mockTransactionRepository) Delete(id uint64) error {
	DeleteWasCalledOnRepository = true
	return m.err
}
