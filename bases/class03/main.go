package main

import (
	"fmt"

	"github.com/GuiTadeu/meli-go/bases/class03/exercises"
)

func main() {

	println(exercises.Tax(49999))  // 0
	println(exercises.Tax(50000))  // 17
	println(exercises.Tax(150000)) // 10

	fmt.Println(exercises.Avg(4, 4, 4, 4))    // 4
	fmt.Println(exercises.Avg(10, 5, 10, 5))  // 7.5
	fmt.Println(exercises.Avg(10, 10, 10, 5)) // 8.75

	avg, err := exercises.Avg(10, 10, -10, 5)
	if err != nil {
		fmt.Println(err) // Error - Número negativo
	} else {
		fmt.Println(avg)
	}

	fmt.Println(exercises.Salary("C", 160)) // 160.000
	fmt.Println(exercises.Salary("B", 160)) // 240.000
	fmt.Println(exercises.Salary("A", 160)) // 480.000

	fmt.Println(exercises.Salary("C", 161)) // 161.000
	fmt.Println(exercises.Salary("B", 161)) // 289.800
	fmt.Println(exercises.Salary("A", 161)) // 724.500

	minFunc, _ := exercises.Stats("minimum")
	minValue := minFunc(2, 4, 6)
	fmt.Println(minValue) // 2

	maxFunc, _ := exercises.Stats("maximum")
	maxValue := maxFunc(2, 4, 6)
	fmt.Println(maxValue) // 6

	avgFunc, _ := exercises.Stats("average")
	avgValue := avgFunc(2, 4, 6)
	fmt.Println(avgValue) // 4

	subFunc, err := exercises.Stats("subtraction")
	subFuncExist := (subFunc != nil)
	fmt.Println(subFuncExist) // False
	fmt.Println(err)          // Operação não encontrada

	var amount float32

	animalDog, err := exercises.Animals("dog")
	amount += animalDog(5) // 50.000

	animalCat, err := exercises.Animals("cat")
	amount += animalCat(5) // 25.0000

	animalHamster, err := exercises.Animals("hamster")
	amount += animalHamster(4) // 1.000

	animalSpider, err := exercises.Animals("spider")
	amount += animalSpider(8) // 1.200

	fmt.Println(amount) // 77.200
}
