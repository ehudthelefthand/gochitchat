package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}
	server.ListenAndServe()
}
