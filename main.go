package main

import (
	"fmt"

	"example.com/hello_go/calc"
)

func main() {
	var numbers = []float64{1, 2, 3, 4, 5}

	fmt.Printf("Array: %v\nAverage: %v\n", numbers, calc.Average(numbers...))
	fmt.Printf("Std Deviation: %v\n", calc.StandardDeviation(numbers...))

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	number := 5
	iterative_fact := calc.IterativeFactorial(number)
	recursive_fact := calc.RecursiveFactorial(number)

	fmt.Printf("Iterative Factorial: %v! = %v\n", number, iterative_fact)
	fmt.Printf("Recursive Factorial: %v! = %v\n", number, recursive_fact)
	return
}
