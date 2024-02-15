package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AddCorsMiddleWare(router *gin.Engine) *gin.Engine {
	originNonRelease, originRelease := "http://localhost:3000", "https://is-in-bible-br.vercel.app"
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if origin == originNonRelease && gin.Mode() != "release" {
				return true
			} else if origin == originRelease && gin.Mode() == "release" {
				return true
			}
			return false
		},
		AllowMethods: []string{"GET", "OPTIONS"},
		MaxAge:       12 * time.Hour,
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))
	return router
}
