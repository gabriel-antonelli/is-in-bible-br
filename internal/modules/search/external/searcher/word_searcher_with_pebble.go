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
	db := config.GetDB()
	val, closer, err := db.Get([]byte(word))
	if err != nil {
		log.Fatalf("error getting key: %v", err)
	}
	err = closer.Close()
	if err != nil {
		log.Fatalf("error closing get: %v", err)
	}
	intVal, err := strconv.Atoi(string(val))
	if err != nil {
		log.Fatalf("error converting string to int: %v", err)
	}
	return intVal
}
