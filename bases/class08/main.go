package main

import (
	"fmt"

	"github.com/GuiTadeu/meli-go/bases/class08/exercises"
)

func main() {

	_, _ = exercises.RegisterCustomer("Jubileu", "123456789", "+55 11 97070-7070", "Rua dos Bobos, Nº 0")
	// Deu ruim! Erro: O arquivo indicado não foi encontrado ou está danificado

	fmt.Println("\nO panic não parou o programa! :D")
}
