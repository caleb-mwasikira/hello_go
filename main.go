package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func genRandomArray(size int) []int {
	nums := []int{}

	for i := 0; i < size; i++ {
		randNum := rand.Intn(10)
		nums = append(nums, randNum)
	}

	return nums
}

func calcSumOfArray(nums []int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

func main() {
	var (
		seq_total  int   = 0
		conc_total int64 = 0
	)

	nums := genRandomArray(1000000)

	start := time.Now()
	seq_total = calcSumOfArray(nums)
	duration := time.Since(start)

	fmt.Printf("sequential total: %v \t in %v\n", seq_total, duration)

	// concurrent soln
	start = time.Now()
	offset := 10000
	wg := sync.WaitGroup{}

	for i := 0; i < len(nums); i += offset {
		begin := i
		end := i + offset
		smaller_arr := nums[begin:end]

		wg.Add(1)
		go func(smaller_arr []int) {
			defer wg.Done()
			small_total := calcSumOfArray(smaller_arr)
			atomic.AddInt64(&conc_total, int64(small_total))
		}(smaller_arr)
	}
	wg.Wait()
	duration = time.Since(start)

	fmt.Printf("concurrent total: %v \t in %v\n", conc_total, duration)
}
