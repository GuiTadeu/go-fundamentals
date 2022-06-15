package exercises

import (
	itrl "github.com/GuiTadeu/meli-go/tests/class03/exercises/internal"
)

type MockService struct {
	repo MockRepository
}

func (serv *MockService) Update(
	id uint64, code string, currency string, value float64,
	issuer string, receiver string, createdAt string,
) (itrl.Transaction, error) {
	return serv.repo.Update(id, code, currency, value, issuer, receiver, createdAt)
}

func (serv *MockService) Delete(id uint64) error {
	return serv.repo.Delete(id)
}