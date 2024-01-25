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
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := routes.SetupRouter()

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

	gracefulShutdown(srv)
}

func gracefulShutdown(srv *http.Server) {
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	log.Println("Closing db")
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
