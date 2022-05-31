package exercises

import (
	"errors"
	"fmt"
	"os"
)

var id int64 = 0
var database []Customer

func genId() *int64 {
	id += 1
	return &id
}

type Customer struct {
	Id int64
	Name string
	Document string
	Phone string
	Address string
}

func RegisterCustomer(name string, document string, phone string, address string) (*Customer, error) {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Printf("Deu ruim! Erro: %s", err)
		}
	}()

	customerId := genId()
	if(customerId == nil) {
		panic("Problema ao gerar o ID...")
	}

	existCustomer, err := os.ReadFile("./customer.txt")

	if(err != nil) {
		panic("O arquivo indicado não foi encontrado ou está danificado")	
	}

	if(existCustomer != nil) {
		return nil, errors.New("Erro no cadastro: Cliente já existente!")
	}

	customerToRegister := Customer{*customerId, name, document, phone, address}
	database = append(database, customerToRegister)
	
	return &customerToRegister, nil
}