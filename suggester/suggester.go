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

// generate all ordered combinations of query of len >= 2
func GenerateCombinations(word string) <-chan string {
	c := make(chan string)

	go func(c chan string) {
		defer close(c)                 // close the channel once we're finished
		nextCombination(c, word, 2, 0) // start with two-letter words
	}(c)

	return c
}

func nextCombination(c chan string, word string, length int, offset int) {
	wordLen := len(word)

	if offset > wordLen-length {
		offset = 0
		length++
	}

	if length > wordLen {
		return
	}

	c <- word[offset : offset+length]
	nextCombination(c, word, length, offset+1)
}

// The above two methods can be replaced with this simple non-recursive method,
// but I wanted to play with channels
/*
func Combinations(word string) []string {
	wordLen := len(word)
	combinations := make([]string, wordLen*(wordLen-1)/2)
	combinationsIdx := 0
	for len := 2; len <= wordLen; len++ {
		for i := 0; i <= wordLen-len; i++ {
			combinations[combinationsIdx] = word[i : i+len]
			combinationsIdx++
		}
	}
	return combinations
}
*/

func main() {
	var err error

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <letters>\n", os.Args[0])
		os.Exit(1)
	}

	// get and validate query string
	query := os.Args[1]
	valid, _ := regexp.MatchString("^[A-Za-z]{2,}$", query)
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
	query = word.Alphabetize(strings.ToUpper(query))
	wordList := make([]word.Word, 0)
	for candidate := range GenerateCombinations(query) {
		words := wordMap[candidate]
		for _, w := range words {
			wordList = append(wordList, word.New(w))
		}
	}

	// display in order of descending point value
	sort.Sort(sort.Reverse(word.ByScore(wordList)))
	for _, word := range wordList {
		fmt.Println(word)
	}
}
