package config

import (
	"log"

	"github.com/cockroachdb/pebble"
)

var pebbleDB *pebble.DB

func GetDB(dbPath string) *pebble.DB {
	if pebbleDB != nil {
		log.Println("pebbleDB is not nil")
		return pebbleDB
	}
	log.Println("pebbleDB is nil")
	if dbPath == "" {
		dbPath = "words-in-the-bible-db"
	}
	db, err := pebble.Open(dbPath, &pebble.Options{})
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}
	pebbleDB = db
	return db
}
