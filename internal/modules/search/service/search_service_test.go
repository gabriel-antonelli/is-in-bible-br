package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWordNormalizer struct {
	mock.Mock
}
type MockWordSearcher struct {
	mock.Mock
}

func (m *MockWordNormalizer) NormalizeWord(word string) string {
	args := m.Called(word)
	return args.String(0)
}

func (m *MockWordSearcher) Total(word string) int {
	args := m.Called(word)
	return args.Int(0)
}

func TestSearchServiceSearch(t *testing.T) {
	wordNormalizer := new(MockWordNormalizer)
	wordSearcher := new(MockWordSearcher)

	wordNormalizer.On("NormalizeWord", "JEsUs").Return("jesus")
	wordNormalizer.On("NormalizeWord", "jesUS").Return("jesus")
	wordNormalizer.On("NormalizeWord", "aMoR").Return("amor")
	wordNormalizer.On("NormalizeWord", "a").Return("a")
	wordSearcher.On("Total", "jesus").Return(42)
	wordSearcher.On("Total", "amor").Return(23)

	searchService := NewSearchService(wordNormalizer, wordSearcher)

	words := "JEsUs+aMoR+a+jesUS"
	results := searchService.Search(words)

	assert.Equal(t, 2, len(results), "Expected two results")
	assert.Equal(t, 42, results["jesus"], "Expected result for 'jesus'")
	assert.Equal(t, 23, results["amor"], "Expected result for 'amor'")
}
