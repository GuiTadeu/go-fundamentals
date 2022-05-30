package exercises

import (
    "log"
    "time"
)

func InsertionSort(numbers []int) {
	start := time.Now()
	var n = len(numbers)

    for i := 1; i < n; i++ {
        j := i
        for j > 0 {
            if numbers[j-1] > numbers[j] {
                numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
            }
            j = j - 1
        }
    }

	log.Printf("InsertionSort, execution time %s\n", time.Since(start))
}

func SelectionSort(numbers []int) {
	start := time.Now()
	var n = len(numbers)

    for i := 0; i < n; i++ {
        var minIdx = i
        for j := i; j < n; j++ {
            if numbers[j] < numbers[minIdx] {
                minIdx = j
            }
        }
        numbers[i], numbers[minIdx] = numbers[minIdx], numbers[i]
    }

	log.Printf("SelectionSort, execution time %s\n", time.Since(start))
}

func BubbleSort(numbers []int) {
	start := time.Now()
	n := len(numbers)
    sorted := false

    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if numbers[i] > numbers[i+1] {
                numbers[i+1], numbers[i] = numbers[i], numbers[i+1]
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        n = n - 1
    }

	log.Printf("BubbleSort, execution time %s\n", time.Since(start))
}