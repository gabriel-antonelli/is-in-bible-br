package validation

func IsValidWord(word string) bool {
	return len(word) > 1 && word != ""
}
