package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config
}

type config struct {
	addr string
}

/*
Mux is good but to group api's third party library would do better
chi - routing library: https://go-chi.io/#/README
*/
// func (app *application) mount() *http.ServeMux {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("GET /health", app.healthCheckHandler)

// 	return mux
// }

// chi implementation
func (app *application) mount() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r

}

func (app *application) run(mux *chi.Mux) error {

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
