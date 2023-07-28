package io

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ReadFileStats(filename string) {
	defer fmt.Println("=========")

	stats, err := os.Stat(filename)
	if err != nil {

		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Filename: %v\n", stats.Name())
	fmt.Printf("Size in bytes: %v\n", stats.Size())
	fmt.Printf("Permissions: %v\n", stats.Mode())
	fmt.Printf("Last Modified: %v\n", stats.ModTime().Format(time.UnixDate))
	fmt.Printf("Is Dir? %v\n", stats.IsDir())
}

func ReadWholeFile(filename string) {
	fmt.Printf("Reading file %v as whole...\n", filename)
	defer fmt.Println("=========")

	buffer, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(buffer))
}

func ReadByLine(filename string) {
	fmt.Printf("Reading file %v line by line...\n", filename)
	defer fmt.Println("=========")

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByWord(filename string) {
	fmt.Printf("Reading file %v word by word...\n", filename)
	defer fmt.Println("=========")

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}