package config

import (
	"log"

	"github.com/cockroachdb/pebble"
)

var pebbleDB *pebble.DB

func GetDB() *pebble.DB {
	if pebbleDB != nil {
		return pebbleDB
	}
	db, err := pebble.Open("words-in-the-bible-db", &pebble.Options{})
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}
	pebbleDB = db
	return db
}
