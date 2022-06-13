package exercises

import "errors"

func Sum(n1, n2 int) int {
	return n1 + n2
}

func Sub(n1, n2 int) int {
	return n1 - n2
}

func Mul(n1, n2 int) int {
	return n1 * n2
}

func Div(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("O denominador não pode ser 0")
	}
	return n1 / n2, nil
}