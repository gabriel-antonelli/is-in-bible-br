package ports

type WordSearcher interface {
	Total(word string) int
}
