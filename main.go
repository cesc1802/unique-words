package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	uniqueWords := FilterUniqueWords(args)

	fmt.Println(uniqueWords)
}

// FilterUniqueWords traverse through the list and return
// a list of unique words appearing in input list
func FilterUniqueWords(words []string) []string {
	uniqueWords := []string{}

	for _, word := range words {
		existed := false
		for _, uniqueWord := range uniqueWords {
			if uniqueWord == word {
				existed = true
			}
		}
		if !existed {
			uniqueWords = append(uniqueWords, word)
		}
	}

	return uniqueWords
}
