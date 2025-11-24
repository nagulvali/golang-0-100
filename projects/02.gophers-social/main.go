package main

import "net/http"

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from web server"))
}

func main() {
	s := &server{addr: ":8080"}

	http.ListenAndServe(":8080", s)
}
