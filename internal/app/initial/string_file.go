package initial

import (
	"log"
	"os"
)

var StringFile string

func FileToString(path string) {
	log.Printf("Loading file %s", path)
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error loading string file: %v", err)
	}
	StringFile = string(file)
}
