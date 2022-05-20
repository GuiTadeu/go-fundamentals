package main

import "github.com/GuiTadeu/meli-go/bases/class02/exercises"

func main() {
	exercises.Letters()

	exercises.Bank(21, true, 13, 100001)  // Não elegível
	exercises.Bank(22, true, 13, 100001)  // Não elegível
	exercises.Bank(23, false, 13, 100001) // Não elegível
	exercises.Bank(23, true, 11, 100001)  // Não elegível
	exercises.Bank(23, true, 12, 100001)  // Não elegível
	exercises.Bank(23, true, 12, 100000)  // Não elegível

	exercises.Bank(23, true, 13, 100000) // Elegível com juros
	exercises.Bank(23, true, 13, 100001) // Elegível sem juros

	exercises.Months(1)  // Janeiro
	exercises.Months(7)  // Julho
	exercises.Months(12) // Dezembro

	exercises.Employees()
}
