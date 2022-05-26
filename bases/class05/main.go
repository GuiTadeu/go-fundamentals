package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/GuiTadeu/meli-go/bases/class05/exercises"
)

func main() {
	exercises.WriteProducts()

	products := exercises.ReadProducts()

	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer writer.Flush()

	fmt.Fprintf(writer, "\n %s\t%s\t%s\t", "ID", "Preço", "Quantidade")

	total := 0.0
	for _, product := range products {

		price, _ := strconv.ParseFloat(product.Price, 32)
		quantity, _ := strconv.ParseFloat(product.Quantity, 32)
		total += price * quantity

		fmt.Fprintf(writer, "\n %s\t%s\t%s\t", product.Id, product.Price, product.Quantity)
	}

	fmt.Fprintf(writer, "\n\t%f\t", total)
	fmt.Fprintf(writer, "\n \t \n")

}
