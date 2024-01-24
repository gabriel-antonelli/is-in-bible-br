package routes

import (
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/app/adapters"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/search/controller"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/search/external/normalizer"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/search/external/searcher"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/modules/search/service"
	"github.com/gin-gonic/gin"
)

func SearchRoutes(router *gin.Engine) {
	router.GET("/search/:word", func(c *gin.Context) {
		wordNormalizer := normalizer.NewWordNormalizerWithText()
		wordSearcher := searcher.NewWordSearcherWithPebble()
		service := service.NewSearchService(wordNormalizer, wordSearcher)
		controller := controller.NewSearchController(service)
		adapters.ControllerAdapter(controller, []string{"word"}, c)
	})
}
