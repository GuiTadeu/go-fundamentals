package exercises

import "errors"

func Animals(name string) (func(quantity uint32) float32, error) {

	const (
		dog     = "dog"
		cat     = "cat"
		hamster = "hamster"
		spider  = "spider"
	)

	switch name {
	case dog:
		return countDog, nil
	case cat:
		return countCat, nil
	case hamster:
		return countHamster, nil
	case spider:
		return countSpider, nil
	default:
		return nil, errors.New("Animal não encontrado!")
	}

}

func countDog(quantity uint32) float32 {
	return float32(10000 * quantity)
}

func countCat(quantity uint32) float32 {
	return float32(5000 * quantity)
}

func countHamster(quantity uint32) float32 {
	return float32(250 * quantity)
}

func countSpider(quantity uint32) float32 {
	return float32(150 * quantity)
}
