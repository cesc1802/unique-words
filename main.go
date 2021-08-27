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

	m := make(map[string]bool)

	for _, word := range words {
		if m[word] {
			continue
		}
		uniqueWords = append(uniqueWords, word)
		m[word] = true
	}

	return uniqueWords
}
