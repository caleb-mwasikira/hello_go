package main

import "fmt"

/*
Write a function to count the frequency of words in a string of text and return a map
of words with their counts. The function should convert the text to lowercase, and
punctuation should be trimmed from words. The strings package contains several help-
ful functions for this task, including Fields, ToLower , and Trim .
Use your function to count the frequency of words in the following passage and then
display the count for any word that occurs more than once:

The string is available in the `data/out_of_the_silent_planet.md` file
*/

func main() {
	filename := "data/out_of_the_silent_planet.md"
	word_count := freqOfWordsInFile(filename)

	for word, count := range word_count {
		// only printing words that have appeared more than once
		// to avoid extremely verbose output to stdout
		if count > 1 {
			fmt.Printf("%v: %v\n", word, count)
		}
	}

}
