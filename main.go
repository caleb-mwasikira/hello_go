package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int64
var delay time.Duration = 100 * time.Millisecond
var mutex = &sync.Mutex{}

func credit(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		mutex.Lock()

		balance += 100
		time.Sleep(delay)
		fmt.Println("After crediting, balance is", balance)

		mutex.Unlock()
	}
}

func debit(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		mutex.Lock()

		balance -= 100
		time.Sleep(delay)
		fmt.Println("After debiting, balance is", balance)

		mutex.Unlock()
	}
}

func main() {
	balance = 200
	wg := sync.WaitGroup{}

	fmt.Println("Initial balance is ", balance)

	wg.Add(1)
	go credit(&wg)

	wg.Add(1)
	go debit(&wg)

	// block until the WaitGroup counter is 0
	wg.Wait()

	fmt.Println("final balance is ", balance)
}
