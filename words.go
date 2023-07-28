package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	s "example.com/hello_go/my_slices"
)

/*
Representation of special characters as []string
*/
func genSpecialCharacters() []string {
	// special character ASCII numbers
	runes := s.Join(
		s.Range(33, 47),
		s.Range(58, 64),
		s.Range(91, 96),
		s.Range(123, 126),
	)

	// convert special character ASCII numbers into characters
	special_chars := make([]string, len(runes))

	for i := range runes {
		special_chars[i] = string(rune(runes[i]))
	}
	return special_chars
}

var special_chars []string = genSpecialCharacters()

/*
Removes all special characters from a string.
*/
func stripSpecialChars(str string) string {
	words := strings.Fields(str)
	parsed_words := make([]string, len(words))

	for i, word := range words {
		parsed_words[i] = strings.Trim(word, strings.Join(special_chars, ""))
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
