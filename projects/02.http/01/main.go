package main

import "net/http"

// changed struct name from server to api/application
type api struct {
	addr string
}

// interface method is same
func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello from web server"))

	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("main page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		default:
			w.Write([]byte("404 page not found"))
		}
	default:
		w.Write([]byte("404 page not found"))
	}
}

func main() {
	// create handler
	api := &api{
		addr: ":8080",
	}

	// use http.Server struct update, instead of passing addr & handler struct as parameters to http.ListenAndServe(addr, handler) direct function
	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}

	srv.ListenAndServe() // calling direct struct method instead of http.ListenAndServe direct package function

}
