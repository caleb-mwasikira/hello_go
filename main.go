package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int64
var delay time.Duration = 100 * time.Millisecond
var mutex = &sync.Mutex{}

func credit() {
	for i := 0; i < 5; i++ {
		mutex.Lock()

		balance += 100
		time.Sleep(delay)
		fmt.Println("After crediting, balance is ", balance)

		mutex.Unlock()
	}
}

func debit() {
	for i := 0; i < 5; i++ {
		mutex.Lock()

		balance -= 100
		time.Sleep(delay)
		fmt.Println("After debiting, balance is ", balance)

		mutex.Unlock()
	}
}

func main() {
	balance = 200
	fmt.Println("Initial balance is ", balance)

	go credit()
	go debit()
	fmt.Scanln()
}
