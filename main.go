package main

import (
	"fmt"
	"time"
)

var balance int64
var delay time.Duration = 100 * time.Millisecond

func credit() {
	for i := 0; i < 5; i++ {
		balance += 100
		time.Sleep(delay)
		fmt.Println("After crediting, balance is ", balance)
	}
}

func debit() {
	for i := 0; i < 5; i++ {
		balance -= 100
		time.Sleep(delay)
		fmt.Println("After debiting, balance is ", balance)
	}
}

func main() {
	balance = 200
	fmt.Println("Initial balance is ", balance)

	go credit()
	go debit()
	fmt.Scanln()
}
