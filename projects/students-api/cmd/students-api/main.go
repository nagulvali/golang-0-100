package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nagulvali/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	
	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to strudents api"))
	})

	// setup server
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}


	slog.Info("Server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()	
		if err != nil {
			log.Fatal("Failed to start server", err)
		}
	} ()

	<- done

	slog.Info("Shutting down the server")

	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown the server.", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown sucessfully")



}
