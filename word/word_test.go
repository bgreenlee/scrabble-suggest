package word

import (
	"sort"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	word := New("hello")

	if word.Word != "HELLO" {
		t.Error("Expected word to be HELLO, got ", word.Word)
	}

	if word.Score != 8 {
		t.Error("Expected score to be 8, got ", word.Score)
	}
}

func TestSortByScore(t *testing.T) {
	words := []Word{
		New("HELLO"),
		New("AA"),
		New("ZAP"),
	}

	sort.Sort(sort.Reverse(ByScore(words)))

	wordList := make([]string, 0)
	for _, word := range words {
		wordList = append(wordList, word.Word)
	}

	if strings.Join(wordList, ",") != "ZAP,HELLO,AA" {
		t.Error("Reverse sort failed. Got ", words)
	}
}

func TestAlphabetize(t *testing.T) {
	sorted := Alphabetize("HELLO")

	if sorted != "EHLLO" {
		t.Error("Expected EHLLO, got ", sorted)
	}
}

func TestFilteredPowerset(t *testing.T) {
	combos := FilteredPowerset("ABC")
	sort.Strings(combos)
	expected := "AB,ABC,AC,BC"
	result := strings.Join(combos, ",")
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
