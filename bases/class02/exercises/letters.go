package exercises

import "fmt"

func Letters() {

	const word = "Zeppelin"
	const size = len(word)

	fmt.Println(size)

	for i := 0; i < size; i++ {
		letter := string(word[i])
		fmt.Print(letter + " ")
	}

	fmt.Println()
}
