package exercises

import "fmt"

func Months(num uint8) {

	months := [12]string{
		"Janeiro",
		"Fevereiro",
		"Março",
		"Abril",
		"Maio",
		"Junho",
		"Julho",
		"Agosto",
		"Setembro",
		"Outubro",
		"Novembro",
		"Dezembro",
	}

	fmt.Println(months[num-1])

}
