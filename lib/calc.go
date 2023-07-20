package lib

import (
	"math"
)

func Average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}

func StandardDeviation(numbers ...float64) float64 {
	// Step 1: Work out the mean
	mean := Average(numbers...)

	// Step 2: For each number: Subtract the mean and square the result
	var squared_diffs []float64
	for _, num := range numbers {
		squared_diffs = append(squared_diffs, math.Pow((num-mean), 2))
	}

	// Step 3: Work out the mean of the squared differences
	mean_squared_diffs := Average(squared_diffs...)

	// Step 4: Square root the result
	return math.Sqrt(mean_squared_diffs)
}

func IterativeFactorial(number int) int {
	factorial := 1
	if number < 1 {
		panic("Cannot compute factorial for negative numbers")
	}

	for i := number; i > 0; i-- {
		factorial *= i
	}
	return factorial
}

func RecursiveFactorial(number int) int {
	if number == 1 {
		return 1
	} else if number < 1 {
		panic("Cannot compute factorial for negative numbers")
	}

	return number * RecursiveFactorial(number-1)
}
