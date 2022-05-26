package exercises

import (
	"fmt"
	"os"
	"strings"
)

type Product struct {
	id       uint64
	price    float32
	quantity uint8
}

type ProductOutput struct {
	Id       string
	Price    string
	Quantity string
}

func WriteProducts() {

	p1 := Product{01, 2000.00, 2}
	p2 := Product{02, 400.00, 10}
	p3 := Product{03, 900.00, 1}

	products := []Product{p1, p2, p3}

	header := []byte("id;price;quantity\n")
	csv := []byte(header)

	for index, p := range products {
		var str string
		if index == (len(products) - 1) {
			str = fmt.Sprintf("%d;%f;%d", p.id, p.price, p.quantity)
		} else {
			str = fmt.Sprintf("%d;%f;%d\n", p.id, p.price, p.quantity)
		}
		csv = append(csv, []byte(str)...)
	}

	os.WriteFile("./products.csv", csv, 0644)
}

func ReadProducts() (outputList []ProductOutput) {

	data, err := os.ReadFile("./products.csv")

	if err != nil {
		panic(err)
	}

	productsCsv := strings.Split(string(data), "\n")

	for index, product := range productsCsv {

		// Remove CSV Header
		if index == 0 {
			continue
		}

		pValues := strings.Split(string(product), ";")

		id := pValues[0]
		price := pValues[1]
		quantity := pValues[2]

		outputList = append(outputList, ProductOutput{id, price, quantity})
	}

	return
}
