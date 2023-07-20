package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello Go!")

	var (
		max       uint64 = 10
		min       uint64 = 1
		str_input string
		number    uint64
	)

	for {
		fmt.Printf("Enter a number between %d-%d: ", min, max)
		fmt.Scanln(&str_input)

		var err error
		number, err = strconv.ParseUint(str_input, 10, 64)
		if err != nil || (number < min || number > max) {
			fmt.Printf("Number MUST be between %d-%d\n", min, max)
			continue
		}

		break
	}

	fmt.Printf("You entered the number: %v\n", number)
	return
}
