package main

import "net/http"

type server struct {
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi from go web server!"))
}

func main() {
	server := &server{}
	http.ListenAndServe(":8080", server)
}
