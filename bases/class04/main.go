package main

import (
	"fmt"

	"github.com/GuiTadeu/meli-go/bases/class04/exercises"
)

func main() {

	arrStudents := [3]exercises.Student{
		exercises.Build("Guilherme", "Tadeu", "123456789", 2022),
		exercises.Build("Guilherme", "Tadeu", "123456789", 2022),
		exercises.Build("Guilherme", "Tadeu", "123456789", 2022),
	}

	fmt.Println(arrStudents)

	p1 := exercises.NewProduct("Headphone", "small", 300.00) // 300
	p2 := exercises.NewProduct("TV", "medium", 2000.00)      // 2060
	p3 := exercises.NewProduct("PS5", "large", 5000.00)      // 7800

	cart := exercises.NewCart()

	cart.Add(p1)
	cart.Add(p2)
	cart.Add(p3)

	fmt.Println(cart.Total()) // 10.160
}
