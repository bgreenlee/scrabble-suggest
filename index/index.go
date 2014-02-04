package index

import (
	"encoding/gob"
	"os"
)

const indexFile = "data/words.idx"

// Write out the index file
func Write(wordMap map[string][]string) (err error) {
	outputFile, err := os.Create(indexFile)
	if err != nil {
		return err
	}

	return gob.NewEncoder(outputFile).Encode(wordMap)
}

func Read() (w map[string][]string, err error) {
	wordMap := make(map[string][]string)
	inputFile, err := os.Open(indexFile)
	if err != nil {
		return nil, err
	}
	err = gob.NewDecoder(inputFile).Decode(&wordMap)
	return wordMap, err
}
