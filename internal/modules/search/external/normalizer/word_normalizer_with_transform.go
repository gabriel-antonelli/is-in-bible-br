package normalizer

import (
	"strings"
	"unicode"

	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/search/ports"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type wordNormalizerWithText struct{}

func NewWordNormalizerWithText() ports.WordNormalizer {
	return &wordNormalizerWithText{}
}

func (normalizer *wordNormalizerWithText) NormalizeWord(word string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	word, _, _ = transform.String(t, word)
	return strings.ToLower(word)
}
