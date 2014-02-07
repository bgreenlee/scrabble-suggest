package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestBuildWordMap(t *testing.T) {
	words := []string{"ate", "eta", "pats", "past", "tea"}
	stringScanner := bufio.NewScanner(strings.NewReader(strings.Join(words, "\n")))
	wordMap := buildWordMap(stringScanner)
	if len(wordMap["AET"]) != 3 && len(wordMap["APST"]) != 2 {
		t.Error("Error building word map. Got ", wordMap)
	}
}
