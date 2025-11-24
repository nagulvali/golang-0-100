package main

import "net/http"

type api struct {
	addr string
}

func (api *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users list"))
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main page"))
	})
	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	srv.ListenAndServe()

}
