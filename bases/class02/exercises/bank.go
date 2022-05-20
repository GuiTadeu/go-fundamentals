package exercises

import "fmt"

func Bank(
	age uint8,
	working bool,
	timeWorkingMonths uint8,
	salary uint32,
) {

	if (age > 22) && (working == true) && (timeWorkingMonths > 12) {

		fmt.Println("Empréstimo elegível")

		if salary > 100000 {
			fmt.Println("Sem Juros")
		} else {
			fmt.Println("Com Juros")
		}
	} else {
		fmt.Println("Empréstimo não elegível")
	}
}
