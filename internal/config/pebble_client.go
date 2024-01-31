package config

import (
	"log"

	"github.com/cockroachdb/pebble"
)

var pebbleDB *pebble.DB

func GetDB(dbPath string) *pebble.DB {
	if pebbleDB != nil {
		return pebbleDB
	}
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

func CloseDB() {
	if pebbleDB != nil {
		err := pebbleDB.Close()
		if err != nil {
			log.Fatalf("error closing db: %v", err)
		}
	}
	pebbleDB = nil
}
