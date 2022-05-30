package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/GuiTadeu/meli-go/bases/class06/exercises"
)

func main() {

	fmt.Println("start main")

	user := exercises.User{
		Name: "Jubileu",
		Age: 22,
		Email: "jubileu@gmail.com",
		Password: "jub123",
	}

	fmt.Println(user) // {Jubileu 22 jubileu@gmail.com jub123}

	user.ChangeName("Pica-pau")
	user.ChangeAge(23)
	user.ChangeEmail("picapau@gmail.com")
	user.ChangePassword("pp123")

	fmt.Println(user) // {Pica-pau 23 picapau@gmail.com pp123}

	c := exercises.Customer{
		Name: "Jack",
		Surname: "White",
		Email: "jack.white@mercadolivre.com",
	}

	p1 := exercises.Build("Headphone", 300.00, 2) // 300
	p2 := exercises.Build("TV", 2000.00, 3)      // 2060
	p3 := exercises.Build("PS5", 5000.00, 4)      // 7800

	exercises.AddProduct(&c, p1)
	exercises.AddProduct(&c, p2)
	exercises.AddProduct(&c, p3)

	fmt.Println(c.Products) // [{Headphone 300 2} {TV 2000 3} {PS5 5000 4}]
	
	exercises.ClearProducts(&c)
	fmt.Println(c.Products) // []

	// Goroutines

	grProducts := []exercises.Product{p1, p2, p3}

	grServices := []exercises.Service{
		{Name:"Instalar Linux", Price:100.0, Minutes:60},
		{Name:"Instalar Windows", Price:100.0, Minutes:120},
	}

	grMaintenance := []exercises.Maintenance{
		{Name:"Atualizar SO", Price:40},
		{Name:"Conectar VPN", Price:50},
	}

	go exercises.TotalProducts(grProducts) // 7300
	go exercises.TotalServices(grServices) // 300
	go exercises.TotalMaintenance(grMaintenance) // 90


	go exercises.Runner("Runner 1")
	go exercises.Runner("Runner 2")
	go exercises.Runner("Runner 3")

	random100 := rand.Perm(100)
	random1000 := rand.Perm(1000)
	random10000 := rand.Perm(10000)

	exercises.InsertionSort(random100) // 6.105µs
	exercises.SelectionSort(random100) // 4.96µs
	exercises.BubbleSort(random100) // 142ns

	exercises.InsertionSort(random1000) // 573.323µs
	exercises.SelectionSort(random1000) // 468.063µs
	exercises.BubbleSort(random1000) // 658ns

	exercises.InsertionSort(random10000) // 43.551835ms
	exercises.SelectionSort(random10000) // 44.149834ms
	exercises.BubbleSort(random10000) // 8.512µs
	
	time.Sleep(10000 * time.Millisecond)

	fmt.Println("end main")
}