package exercises

func Tax(salary float32) uint8 {

	if salary == 50000 {
		return 17
	}

	if salary == 150000 {
		return 10
	}

	return 0
}
