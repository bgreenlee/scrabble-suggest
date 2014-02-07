package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestFindWords(t *testing.T) {
	wordMap := map[string][]string{
		"ABT": []string{"BAT", "TAB"},
		"APT": []string{"APT", "PAT", "TAP"},
	}

	words := findWords("APTXB", wordMap)
	fmt.Println(words)
	wordList := make([]string, 0)
	for _, word := range words {
		wordList = append(wordList, word.Word)
	}
	sort.Strings(wordList)

	result := strings.Join(wordList, ",")
	expected := "APT,BAT,PAT,TAB,TAP"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
