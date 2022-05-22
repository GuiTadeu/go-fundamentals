package exercises

import "errors"

func Stats(operation string) (func(...float32) float32, error) {

	const (
		minimum = "minimum"
		maximum = "maximum"
		average = "average"
	)

	switch operation {
	case minimum:
		return minValue, nil
	case maximum:
		return maxValue, nil
	case average:
		return avgValue, nil
	default:
		return nil, errors.New("Operação não encontrada!")
	}
}

func minValue(values ...float32) (minValue float32) {
	minValue = values[0]

	for i := 1; i < len(values); i++ {
		if values[i] < minValue {
			minValue = values[i]
		}
	}

	return
}

func maxValue(values ...float32) (maxValue float32) {
	maxValue = 0.0

	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return
}

func avgValue(values ...float32) float32 {
	var count, sum float32

	for _, value := range values {
		sum += value
		count++
	}

	return sum / count
}
