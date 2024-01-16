package main

import (
	"log"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

func main() {
	r := searchWord()
	r.Run()
}

func searchWord() *gin.Engine {
	r := gin.Default()
	r.GET("search/:word", func(c *gin.Context) {
		word := c.Param("word")
		log.Printf("Searching for: %s", word)
		total := totalNumberIndexed(word)
		c.JSON(200, gin.H{
			"found": total > 0,
			"total": total,
		})
	})
	return r
}

func totalNumberIndexed(word string) int {
	file, err := os.ReadFile("biblia_normalized.txt")
	if err != nil {
		log.Fatal(err)
	}
	stringFile := string(file)
	regex := regexp.MustCompile(`(?i)\b` + word + `\b`)
	matches := regex.FindAllString(stringFile, -1)
	return len(matches)
}
