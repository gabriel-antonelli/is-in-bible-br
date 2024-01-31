package searcher

import (
	"os"
	"testing"

	"github.com/cockroachdb/pebble"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/config"
	"github.com/stretchr/testify/assert"
)

var (
	testDB  *pebble.DB
	tempDir string
)

func getDB(t *testing.T) *pebble.DB {
	if testDB != nil {
		return testDB
	}
	tempDir = t.TempDir()
	return config.GetDB(tempDir)
}

func teardown(t *testing.T) {
	config.CloseDB()
	err := os.RemoveAll(tempDir)
	if err != nil {
		t.Fatalf("error removing temp dir: %v", err)
	}
}

func TestWordSearcherStringReturned(t *testing.T) {
	err := getDB(t).Set([]byte("jesus"), []byte("1075"), nil)
	if err != nil {
		t.Errorf("error setting key-value: %v", err)
	}

	searcher := NewWordSearcherWithPebble()

	total := searcher.Total("jesus")

	assert.Equal(t, 1075, total, "Total is expected to be more than 0")
}

func TestWordSearcherKeyNotFound(t *testing.T) {
	err := getDB(t).Delete([]byte("jesus"), nil)
	if err != nil {
		t.Errorf("error deleting key: %v", err)
	}
	searcher := NewWordSearcherWithPebble()

	total := searcher.Total("jesus")

	assert.Equal(t, 0, total, "Total is expected to be 0")
}

func TestWordSearcherCannotConvertToInt(t *testing.T) {
	defer teardown(t)
	err := getDB(t).Set([]byte("test"), []byte("word"), nil)
	if err != nil {
		t.Errorf("error setting key-value: %v", err)
	}

	searcher := NewWordSearcherWithPebble()

	total := searcher.Total("test")

	assert.Equal(t, 0, total, "Total is expected to be 0")
}
