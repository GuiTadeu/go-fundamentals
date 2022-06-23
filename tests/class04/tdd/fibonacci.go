package tdd

func Fibonacci(quantity uint8) (sequence []uint8) {

	var n1 uint8 = 1
	var n2 uint8 = 1

	sequence = []uint8{n1, n2}

	for i := 2; i < int(quantity); i++ {
		insertElement := n1 + n2
		n1 = n2
		n2 = insertElement

		sequence = append(sequence, insertElement)
	}

	return
}