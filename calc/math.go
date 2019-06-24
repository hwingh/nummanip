package calc

import "errors"
import "github.com/fatih/color"

// Add return sum of two integers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		errorMessages := color.RedString("provide more than 2 numbers")
		return sum, errors.New(errorMessages)
	} else {
		for _, num := range numbers {
			sum = sum + num
		}

		return sum, nil
	}

}
