package ports

type WordNormalizer interface {
	NormalizeWord(word string) string
}
