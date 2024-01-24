package service

import (
	"log"
	"strings"

	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/search/ports"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/search/validation"
)

type SearchService interface {
	Search(words string) map[string]int
}

type searchService struct {
	wordNormalizer ports.WordNormalizer
	wordSearcher   ports.WordSearcher
}

func NewSearchService(wordNormalizer ports.WordNormalizer, wordSearcher ports.WordSearcher) SearchService {
	return &searchService{wordNormalizer, wordSearcher}
}

func (c *searchService) Search(words string) map[string]int {
	results := make(map[string]int)
	for _, word := range strings.Split(words, "+") {
		if !validation.IsValidWord(word) {
			continue
		}

		word = c.wordNormalizer.NormalizeWord(word)
		_, ok := results[word]
		if ok {
			continue
		}

		log.Printf("Searching for %s\n", word)
		results[word] = c.wordSearcher.Total(word)
	}
	return results
}
