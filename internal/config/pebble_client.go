package config

import (
	"log"

	"github.com/cockroachdb/pebble"
)

var db *pebble.DB

func GetDB() *pebble.DB {
	if db != nil {
		return db
	}
	db, err := pebble.Open("words-in-the-bible-db", &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
