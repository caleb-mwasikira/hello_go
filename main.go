package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
)

func sum(arr []int, ch chan int) {
	total := 0

	for _, val := range arr {
		total += val
	}

	// sending partial sum over channel
	ch <- total
}

/* generates a random array of n elements from 0 - max */
func randomArray(n int, max int) ([]int, error) {
	arr := []int{}

	if n <= 0 || max <= 0 {
		return arr, errors.New("invalid number of elements or max value")
	}

	for i := 0; i < n; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			return []int{}, errors.New("failed to generate random array of numbers")
		}

		arr = append(arr, int(nBig.Int64()))
	}
	return arr, nil
}

/*
splits an array into groups with each group having n elements

for example:

	arr := [a, b, c, d, e, f, g, h, i, j]
	splitArray(arr, 5)
	returns [a, b, c, d, e] and [f, g, h, i, j]
*/
func splitArray(arr []int, n int) [][]int {
	arrays := [][]int{}
	max_index := len(arr) - 1
	parts := math.Round(float64(len(arr)/n)) + 1

	var i int
	for i = 0; i < int(parts); i++ {
		start, stop := i*n, (i+1)*n

		var slice []int
		if stop > max_index {
			slice = arr[start:]
		} else {
			slice = arr[start:stop]
		}
		arrays = append(arrays, slice)
	}

	return arrays
}

func main() {
	random_array, err := randomArray(10, 100)
	if err != nil {
		log.Fatal(err)
	}

	// how would we go about finding the sum of the above random array?
	// one way would be to iterate over each element and add them up sequentially.
	// but as the array grows, our solution takes longer to complete.
	// a solution to reduce the time complexity of our program is to divide the array
	// down into multiple sub-arrays, compute the partial sums of sub arrays on
	// different threads and add the partial sums to get a final sum value

	// split array into sub-arrays with n elements each
	arrays := splitArray(random_array, 6)

	// create channel for sum() goroutines to communicate in
	sum_ch := make(chan int)

	// launch goroutines to calculate sum
	for _, array := range arrays {
		go sum(array, sum_ch)
	}

	// capture messages sent over the channel
	total := 0
	for i := 0; i < len(arrays); i++ {
		partial_sum := <-sum_ch
		total += partial_sum
	}

	fmt.Printf("sum of random array %v is %v\n", random_array, total)
}
