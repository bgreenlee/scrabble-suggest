package main

import (
	"fmt"
	"footle.org/scrabble/index"
	"footle.org/scrabble/word"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// findWords, given a query string and a wordMap index, finds all the words
// in the index that can be made with the letters in the query string
func findWords(query string, wordMap map[string][]string) []word.Word {
	query = word.Alphabetize(strings.ToUpper(query))
	wordList := make([]word.Word, 0)
	for _, candidate := range word.FilteredPowerset(query) {
		words := wordMap[candidate]
		for _, w := range words {
			wordList = append(wordList, word.New(w))
		}
	}
	return wordList
}

func main() {
	var err error

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <letters>\n", os.Args[0])
		os.Exit(1)
	}

	// get and validate query string
	query := os.Args[1]
	valid, _ := regexp.MatchString("^[A-Za-z_]{2,}$", query)
	if !valid {
		fmt.Printf("Error: you must provide 2 or more letters, A-Z\n")
		os.Exit(1)
	}

	// read index
	wordMap, err := index.Read()
	if err != nil {
		log.Fatal("Error reading index:", err)
	}

	// find possible words
	wordList := findWords(query, wordMap)

	// display in order of descending point value
	sort.Sort(sort.Reverse(word.ByScore(wordList)))
	for _, word := range wordList {
		fmt.Println(word)
	}
}
