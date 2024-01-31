package normalizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var normalizer = NewWordNormalizerWithText()

func TestWordNormalizerWithTextNormalizeWordAccents(t *testing.T) {
	wordWithAccents := "Café"
	normalizedWord := normalizer.NormalizeWord(wordWithAccents)
	assert.Equal(t, "cafe", normalizedWord, "Expected word to be normalized without accents")
}

func TestWordNormalizerWithTextNormalizeWordRemainUnchanged(t *testing.T) {
	wordWithoutAccents := "hello"
	normalizedWord := normalizer.NormalizeWord(wordWithoutAccents)
	assert.Equal(t, "hello", normalizedWord, "Expected word to remain unchanged")
}

func TestWordNormalizerWithTextNormalizeWordToLowerCase(t *testing.T) {
	wordMixedCase := "GoLang"
	normalizedWord := normalizer.NormalizeWord(wordMixedCase)
	assert.Equal(t, "golang", normalizedWord, "Expected word to be normalized to lowercase")
}
