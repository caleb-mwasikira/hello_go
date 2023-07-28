package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"example.com/hello_go/lib"
)

/*
Removes all special characters from a string.
*/
func stripSpecialChars(str string) string {
	special_char_runes := lib.ConcatMultipleSlices[int32]([][]int32{
		lib.MakeRange[int32](33, 47), lib.MakeRange[int32](58, 64),
		lib.MakeRange[int32](91, 96), lib.MakeRange[int32](123, 126),
	})
	special_chars := strings.Join(lib.RunesToStrings(special_char_runes), "")
	words := strings.Fields(str)
	parsed_words := make([]string, len(words))

	for i, word := range words {
		parsed_words[i] = strings.Trim(word, special_chars)
	}

	return strings.Join(parsed_words, " ")
}

/* Counts the frequency of words in a given file */
func freqOfWordsInFile(filename string) map[string]int {
	word_count := make(map[string]int)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file %v: %v\n", filename, err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())

		// Here we are going to strip special characters from
		// our line because without this the words 'wood' and 'wood,'
		// would count as two different words instead of one
		parsed_line := stripSpecialChars(line)
		words := strings.Fields(parsed_line)

		for _, word := range words {
			word_count[word]++
		}
	}
	return word_count
}
