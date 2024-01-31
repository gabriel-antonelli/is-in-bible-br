package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func close(t *testing.T, dbPath string) {
	t.Cleanup(func() {
		CloseDB()
		err := os.RemoveAll(dbPath)
		if err != nil {
			t.Fatalf("error removing temp dir: %v", err)
		}
	})
}

func TestGetDB(t *testing.T) {
	dbPath := t.TempDir()
	close(t, dbPath)

	db := GetDB(dbPath)

	assert.NotNil(t, db, "Expected a non-nil pebble.DB instance")
	assert.DirExists(t, dbPath, "Expected temp dir to be created")
	assertDirNotEmpty(t, dbPath)
}

func TestGetDBEmptyString(t *testing.T) {
	close(t, "words-in-the-bible-db")

	db := GetDB("")

	assert.NotNil(t, db, "Expected a non-nil pebble.DB instance")
	assert.DirExists(t, "words-in-the-bible-db", "Expected default dir to be created")
	assertDirNotEmpty(t, "words-in-the-bible-db")
}

func TestCloseDB(t *testing.T) {
	dbPath := t.TempDir()
	close(t, dbPath)

	db := GetDB(dbPath)

	assert.NotNil(t, db, "Expected a non-nil pebble.DB instance")
	assert.DirExists(t, dbPath, "Expected temp dir to be created")
	assertDirNotEmpty(t, dbPath)

	CloseDB()
	db = GetDB("")
	close(t, "words-in-the-bible-db")

	assert.NotNil(t, db, "Expected a non-nil pebble.DB instance")
	assert.DirExists(t, "words-in-the-bible-db", "Expected default dir to be created")
	assertDirNotEmpty(t, "words-in-the-bible-db")
}

func assertDirNotEmpty(t *testing.T, dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		t.Fatalf("error reading directory %s: %v", dirPath, err)
	}

	assert.NotEmpty(t, entries, "Expected directory %s to not be empty", dirPath)
}
