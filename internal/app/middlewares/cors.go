package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AddCorsMiddleWare(router *gin.Engine) *gin.Engine {
	router.Use(cors.Default())
	return router
}
