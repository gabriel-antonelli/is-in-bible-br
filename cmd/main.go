package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/app/routes"
	"github.com/gabriel-antonelli/is-in-the-bible-br/internal/config"
	"github.com/gin-gonic/gin"
)

var path = "biblia_normalized.txt"

func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.SearchRoutes(router)
	return router
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := setupRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	err := config.GetDB().Close()
	if err != nil {
		log.Println("Error closing db: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
