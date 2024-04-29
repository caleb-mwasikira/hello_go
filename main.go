package main

import (
	"fmt"
	"sync"
)

// --- send data into a channel ---
func sendData(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Hello World"
	fmt.Println("data sent into channel")
}

// -- retrieve data from a channel ---
func getData(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("goroutine x says %v\n", <-ch)
}

func main() {
	// create channel
	ch := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go sendData(ch, &wg)

	wg.Add(1)
	go getData(ch, &wg)

	wg.Wait() // block until the WaitGroup goroutine counter equals 0
}
