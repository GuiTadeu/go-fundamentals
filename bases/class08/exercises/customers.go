package exercises

import "os"

func ReadCustomers() {

	_, err := os.ReadFile("./customers.txt")

	if err != nil {
		panic("O arquivo indicado não foi encontrado ou está danificado")
	}
	
}