package searcher

import (
	"fmt"
	"regexp"
)

type WordSearcherWithRegex struct{}

func NewWordSearcherWithRegex() *WordSearcherWithRegex {
	return &WordSearcherWithRegex{}
}

func (w *WordSearcherWithRegex) Total(word string, stringFile string) int {
	wordRegex := regexp.MustCompile(fmt.Sprintf(`(?i)\b%s\b`, word))
	matches := wordRegex.FindAllString(stringFile, -1)
	return len(matches)
}
