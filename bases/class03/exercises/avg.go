package exercises

import "errors"

func Avg(grades ...float32) (float32, error) {

	var count, sum float32

	for _, grade := range grades {

		if grade < 0 {
			return 0, errors.New("Número negativo")
		}

		sum += grade
		count += 1
	}

	return sum / count, nil
}
