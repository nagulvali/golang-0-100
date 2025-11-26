package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", app.healthCheckHandler)

	return mux
}

func (app *application) run(mux *http.ServeMux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // 30 secs is optimal value can be changed based on research
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	// replace log.printf with custom logging
	log.Printf("server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}
