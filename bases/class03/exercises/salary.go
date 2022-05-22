package exercises

import "errors"

func Salary(category string, hours int16) (salary float64, err error) {

	switch category {
	case "C":
		return calcPerHour(1000, hours), nil
	case "B":
		salary := calcPerHour(1500, hours)
		salary += calcPlus(salary, hours, 0.2)
		return salary, nil
	case "A":
		salary := calcPerHour(3000, hours)
		salary += calcPlus(salary, hours, 0.5)
		return salary, nil
	default:
		return -1, errors.New("Categoria não encontrada")
	}
}

func calcPerHour(salary float64, hours int16) float64 {
	return salary * float64(hours)
}

func calcPlus(salary float64, hours int16, percentage float64) float64 {
	if hours >= 160 {
		return salary * percentage
	}
	return salary
}
