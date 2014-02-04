/*
Implements an indexer that takes a Scrabble word list and produces an index
for use by the suggester program.
*/
package main

import (
	"bufio"
	"fmt"
	"footle.org/scrabble/index"
	"footle.org/scrabble/word"
	"log"
	"os"
	"strings"
)

func main() {
	var err error

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <word file>\n", os.Args[0])
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

	// read in word list and store each word in a map where the keys are the
	// words with letters in alphabetical order
	// so you get:
	// AET -> [ATE, ETA, TAE, TEA],
	// APST -> [PATS, PAST, TAPS, SPAT],
	wordMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := strings.ToUpper(scanner.Text())
		sortedWord := word.Alphabetize(token)
		wordMap[sortedWord] = append(wordMap[sortedWord], token)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal("Error reading input file:", err)
	}

	err = index.Write(wordMap)
	if err != nil {
		log.Fatal("Error writing index file:", err)
	}
}
