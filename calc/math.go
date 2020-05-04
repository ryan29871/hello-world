package calc

import "errors"

// Add - return the sum of 2 or more integers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("Provide more than 2 numbers")
	}
	for _, num := range numbers {
		sum = sum + num
	}

	return sum, nil
}
