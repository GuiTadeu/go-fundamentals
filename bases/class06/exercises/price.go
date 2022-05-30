package exercises

import "time"

type Service struct {
	Name string
	Price float32
	Minutes uint8
}

type Maintenance struct {
	Name string
	Price float32
}

func TotalProducts(products []Product) (total float32) {

	println("start TotalProducts")
	for _, product := range products {
		total += product.price * float32(product.quantity)
	}

	time.Sleep(500 * time.Millisecond)
	println("end TotalProducts")
	return
}

func TotalServices(services []Service) (total float32) {

	println("start TotalServices")
	for _, service := range services {
		minutes := service.Minutes
		price := service.Price

		if minutes <= 30 {
			total = price / 2
		} else {
			priceForMinute := price / 60
			total = priceForMinute * float32(minutes)
		}
	}

	time.Sleep(500 * time.Millisecond)
	println("end TotalServices")

	return
}

func TotalMaintenance(maintenances []Maintenance) (total float32) {

	println("start TotalMaintenance")
	for _, maintenance := range maintenances {
		total += maintenance.Price
	}

	time.Sleep(500 * time.Millisecond)
	println("end TotalMaintenance")
	return
}

// Funny :D
func Runner(name string) {
	for i := 0; i < 50; i++ {
		println(name, i)
		time.Sleep((time.Duration(i) * 1000) * time.Millisecond)
	}
}