package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DownloadResult struct {
	url     string
	success bool
}

func downloadFile(url string) bool {
	fmt.Printf("downloading file %v...\n", url)

	// simulate doing some work for n seconds
	time.Sleep(5 * time.Second)

	result := rand.Float32() < 0.5
	return result
}

func main() {
	var (
		start    time.Time
		duration time.Duration
	)

	urls := []string{
		"https://wikipedia.org",
		"https://www.google.com",
		"https://www.w3schools.com",
	}

	// // sequential downloading of files
	// start = time.Now()
	// for _, url := range urls {
	// 	downloadFile(url)
	// 	fmt.Println()
	// }
	// duration = time.Since(start)
	// fmt.Printf("sequential downloading files took %v\n", duration)

	wg := sync.WaitGroup{}
	start = time.Now()
	res_chan := make(chan DownloadResult)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			result := downloadFile(url)
			res_chan <- DownloadResult{
				url:     url,
				success: result,
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(res_chan)
	}()

	for result := range res_chan {
		if result.success {
			fmt.Printf("download %v successful\n", result.url)
		} else {
			fmt.Printf("download %v failed\n", result.url)
		}
	}

	duration = time.Since(start)
	fmt.Printf("concurrent downloading files took %v\n", duration)
}
