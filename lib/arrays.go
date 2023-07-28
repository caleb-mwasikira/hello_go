package lib

import "fmt"

type number interface {
	int | int32 | int64 | float32 | float64
}

func MakeRange[T number](min T, max T) []T {
	var _range []T

	for i := min; i <= max; i++ {
		_range = append(_range, i)
	}
	return _range
}

func ConcatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int
	for _, slice := range slices {
		totalLen += len(slice)
	}

	result := make([]T, totalLen)

	var i int
	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

/* Converts an array of runes into a array of strings */
func RunesToStrings(runes []int32) []string {
	slice := make([]string, len(runes))

	for i := range slice {
		slice[i] = string(rune(runes[i]))
	}
	return slice
}

/* Joins a array of numbers into a string based on the given separator */
func JoinSlice[T number](array []T, sep string) string {
	arrayLen := len(array)
	var str string

	for index, item := range array {
		if index == arrayLen-1 {
			sep = ""
		}
		str += fmt.Sprintf("%v%v", item, sep)
	}
	return str
}

/* Returns a map of the number of times each element appears in an array */
func freqArray[T number](array []T) map[T]int {
	frequencies := make(map[T]int)

	for _, item := range array {
		frequencies[item]++
	}
	return frequencies
}
