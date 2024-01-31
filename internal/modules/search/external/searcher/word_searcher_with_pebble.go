package searcher

import (
	"log"
	"strconv"

	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/config"
)

type WordSearcherWithPebble struct{}

func NewWordSearcherWithPebble() *WordSearcherWithPebble {
	return &WordSearcherWithPebble{}
}

func (w *WordSearcherWithPebble) Total(word string) int {
	db := config.GetDB("")
	val, closer, err := db.Get([]byte(word))
	if err != nil {
		log.Printf("error getting key '%s': %v", word, err)
		return 0
	}
	err = closer.Close()
	if err != nil {
		log.Printf("error closing get: %v", err)
	}
	intVal, err := strconv.Atoi(string(val))
	if err != nil {
		log.Printf("error converting string to int: %v", err)
	}
	log.Printf("Total for word '%s': %d", word, intVal)
	return intVal
}
