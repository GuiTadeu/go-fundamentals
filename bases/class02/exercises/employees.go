package exercises

import "fmt"

func Employees() {

	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	fmt.Println(employees)
	fmt.Println(employees["Benjamin"])

	var countAgeBt20 = 0
	for _, age := range employees {
		if age > 21 {
			countAgeBt20++
		}
	}

	fmt.Println("Há " + fmt.Sprint(countAgeBt20) + " funcionários com mais de 21 anos")

	employees["Frederico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
