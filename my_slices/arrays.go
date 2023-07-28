package my_slices

type number interface {
	int | int32 | int64 | float32 | float64
}

/*
Creates an array/slice with range starting from min...max
(max is inclusive)
For example:

	Range(0, 5) -> [0, 1, 2, 3, 4, 5]
*/
func Range[T number](min T, max T) []T {
	var _range []T

	for i := min; i <= max; i++ {
		_range = append(_range, i)
	}
	return _range
}

/*
Joins multiple slices into one
*/
func Join[T any](slices ...[]T) []T {
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
