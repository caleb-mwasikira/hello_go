package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var balance int64
var delay time.Duration = 100 * time.Millisecond

func credit() {
	for i := 0; i < 5; i++ {
		atomic.AddInt64(&balance, 100)
		time.Sleep(delay)
	}
}

func debit() {
	for i := 0; i < 5; i++ {
		atomic.AddInt64(&balance, -100)
		time.Sleep(delay)
	}
}

func main() {
	balance = 200
	fmt.Println("Initial balance is ", balance)

	go credit()
	go debit()
	fmt.Scanln()

	fmt.Println("final balance is ", balance)
}
