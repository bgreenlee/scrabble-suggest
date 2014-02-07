package index

import (
	"io/ioutil"
	"os"
	"testing"
)

func setUp() {
	// overwrite the default IndexFile with a temp file
	file, err := ioutil.TempFile("", "words")
	if err != nil {
		panic(err)
	}
	IndexFile = file.Name()
}

func tearDown() {
	err := os.Remove(IndexFile)
	if err != nil {
		panic(err)
	}
}

func TestWriteReadIndex(t *testing.T) {
	var err error

	setUp()

	wordMap := map[string][]string{
		"abt": []string{"bat", "tab"},
		"apt": []string{"apt", "pat", "tap"},
	}

	err = Write(wordMap)
	if err != nil {
		t.Error(err)
	}

	newWordMap, err := Read()
	if err != nil {
		t.Error(err)
	}

	if newWordMap["apt"][2] != "tap" {
		t.Error("Map did not deserialize correctly: ", wordMap)
	}

	tearDown()
}
