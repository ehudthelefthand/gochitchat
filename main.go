package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}
	server.ListenAndServe()
}
