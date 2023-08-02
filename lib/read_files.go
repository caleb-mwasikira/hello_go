package lib

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func ReadFileStats(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Filename: %v\n", stats.Name())
	fmt.Printf("Last Modified: %v\n", stats.ModTime().Format(time.UnixDate))
	fmt.Printf("Is Dir? %v\n", stats.IsDir())
}

func ReadWholeFile(filename string) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(buffer))
}

func ReadByLine(filename string) {
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

func ReadWholeCSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.LazyQuotes = true
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadCSVFileByLine(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]string
	csvReader := csv.NewReader(file)
	csvReader.LazyQuotes = true
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			// We reached the end of file
			break
		}
		if err != nil {
			// If there is an error reading a single line, we skip it
			// and continue to the next line
			log.Default().Printf("Failed to read line on csv file: %v\n", err)
			continue
		}
		data = append(data, record)
	}
	return data, nil
}
