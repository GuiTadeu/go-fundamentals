package exercises

import (
	"errors"
	"fmt"
)

// Error Interface Method
func TaxErrorInterface(salary int) (bool, *MyCustomError) {

	if salary < 15000 {
		return false, &MyCustomError{400, "O salário digitado não está dentro do valor mínimo"}
	}

	return true, nil
}

type MyCustomError struct {
	status int
	msg    string
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

// New Function (pkg errors)
func TaxErrorsNew(salary int) (bool, error) {

	if salary < 15000 {
		return false, errors.New("O salário digitado não está dentro do valor mínimo")
	}

	return true, nil
}

// Errorf Function (pkg fmt)
func TaxErrorf(salary int) (bool, error) {

	if salary < 15000 {
		return false, fmt.Errorf("O mínimo tributável é 15.000 e o salário informado é: %d", salary)
	}

	return true, nil
}

// Complete 
func TaxComplete(hours int, valuePerHour float32) (salary float32, err error) {
	salary = float32(hours) * valuePerHour

	// New Function (pkg errors)
	if hours < 80 {
		return salary, errors.New("O trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}

	// Errorf Function (pkg fmt)
	if salary < 7500 {
		return salary, fmt.Errorf("O mínimo tributável é 15.000 e o salário informado é: %f", salary)
	}

	// Error Interface Method
	if salary < 15000 {
		return salary, &MyCustomError{400, "O salário digitado não está dentro do valor mínimo"}
	}

	if salary >= 20000 {
		salary -= (salary * 0.1)
	}

	return
}
