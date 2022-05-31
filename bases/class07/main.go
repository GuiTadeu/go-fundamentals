package main

import (
	"fmt"

	"github.com/GuiTadeu/meli-go/bases/class07/exercises"
)

func main() {
	
	salary, _ := exercises.TaxComplete(80, 187.5)
	fmt.Println(salary) // 15000

	_, err := exercises.TaxComplete(79, 187.5)
	fmt.Println(err) // O trabalhador não pode ter trabalhado menos de 80 horas por mês

	_, err = exercises.TaxComplete(80, 93.7375)
	fmt.Println(err) // O mínimo tributável é 15.000 e o salário informado é: 7499

	_, err = exercises.TaxComplete(80, 93.7375)
	fmt.Println(err) // O mínimo tributável é 15.000 e o salário informado é: 7499

	_, err = exercises.TaxComplete(80, 187.4875)
	fmt.Println(err) // 400 - O salário digitado não está dentro do valor mínimo

	salary, _ = exercises.TaxComplete(80, 250)
	fmt.Println(salary) // 180000
}
