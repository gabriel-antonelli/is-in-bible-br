package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

var stringFile string
var wordRegex *regexp.Regexp

func init() {
	file, err := os.ReadFile("biblia_normalized.txt")
	if err != nil {
		log.Fatal(err)
	}
	stringFile = string(file)
}

func main() {
	r := searchWord()
	r.Run()
}

func searchWord() *gin.Engine {
	r := gin.Default()
	r.GET("search/:word", func(c *gin.Context) {
		word := c.Param("word")
		wordRegex = regexp.MustCompile(fmt.Sprintf(`(?i)\b%s\b`, word))
		total := totalNumberIndexed()
		c.JSON(200, gin.H{
			"found": total > 0,
			"total": total,
		})
	})
	return r
}

func totalNumberIndexed() int {
	matches := wordRegex.FindAllString(stringFile, -1)
	return len(matches)
}
