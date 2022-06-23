package class04

type mockTransactionService struct {
	result any
	err    error
}

func (m mockTransactionService) Update(
	id uint64, code string, currency string, value float64,
	issuer string, receiver string, createdAt string,
) (Transaction, error) {
	if m.err != nil {
		return Transaction{}, m.err
	}
	return m.result.(Transaction), nil
}

func (m mockTransactionService) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}