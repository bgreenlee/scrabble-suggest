package index

import (
	"encoding/gob"
	"os"
)

var IndexFile = "data/words.idx"

// Write serializes the word map to an index file
func Write(wordMap map[string][]string) (err error) {
	outputFile, err := os.Create(IndexFile)
	if err != nil {
		return err
	}

	return gob.NewEncoder(outputFile).Encode(wordMap)
}

// Read deserializes the index file into a word map
func Read() (w map[string][]string, err error) {
	wordMap := make(map[string][]string)
	inputFile, err := os.Open(IndexFile)
	if err != nil {
		return nil, err
	}
	err = gob.NewDecoder(inputFile).Decode(&wordMap)
	return wordMap, err
}
