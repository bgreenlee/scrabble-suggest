package word

import (
	"fmt"
	"sort"
	"strings"
)

type Word struct {
	Word  string
	Score int
}

func New(word string) Word {
	w := Word{
		Word:  strings.ToUpper(word),
		Score: calculateScore(word),
	}

	return w
}

func (w Word) String() string {
	return fmt.Sprintf("%s (%d)", w.Word, w.Score)
}

func calculateScore(word string) int {
	value := map[rune]int{
		'A': 1, 'E': 1, 'I': 1, 'L': 1, 'N': 1, 'O': 1, 'R': 1, 'S': 1, 'T': 1, 'U': 1,
		'D': 2, 'G': 2,
		'B': 3, 'C': 3, 'M': 3, 'P': 3,
		'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
		'K': 5,
		'J': 8, 'X': 8,
		'Q': 10, 'Z': 10,
	}

	var score int
	for _, char := range strings.ToUpper(word) {
		score += value[char]
	}

	return score
}

// rearrange a word so the letters are in alphabetical order
func Alphabetize(word string) string {
	chars := ByChar(word)
	sort.Sort(chars)
	return string(chars)
}

// ByScore implements sort.Interface for []Word based on the Score field
type ByScore []Word

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

// ByChar implements a sort interface for ASCII characters represented by bytes
type ByChar []byte

func (a ByChar) Len() int           { return len(a) }
func (a ByChar) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByChar) Less(i, j int) bool { return a[i] < a[j] }
