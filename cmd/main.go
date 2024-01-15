package main

import (
	"log"

	"github.com/blevesearch/bleve/v2"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal"
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
	index, _ := bleve.Open("biblia.bleve")
	defer index.Close()

	word = internal.NormalizeText(word)

	query := bleve.NewMatchQuery(word)
	search := bleve.NewSearchRequest(query)
	searchResults, _ := index.Search(search)

	if len(searchResults.Hits) > 0 {
		log.Printf("Found %d results for: %s", searchResults.Total, word)
		return searchResults.Hits[0].Size()
	}
	log.Printf("No results found for: %s", word)
	return 0
}
