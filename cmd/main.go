package main

import (
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/app/initial"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/app/routes"
	"github.com/gin-gonic/gin"
)

var path = "biblia_normalized.txt"

func init() {
	initial.FileToString(path)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.SearchRoutes(router)
	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
