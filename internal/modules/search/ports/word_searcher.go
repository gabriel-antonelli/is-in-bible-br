package ports

type WordSearcher interface {
	Total(word string, stringFile string) int
}
