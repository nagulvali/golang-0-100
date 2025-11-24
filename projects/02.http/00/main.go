/*
Basic http server
*/
package main

import "net/http"

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello from web server"))
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Main page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		default:
			w.Write([]byte("404 page not found"))
			return
		}
	default:
		w.Write([]byte("404 page not found"))
		return
	}
}

func main() {
	s := &server{addr: ":8080"}

	http.ListenAndServe(s.addr, s)
}
