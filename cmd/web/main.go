package main

import (
	"log"
	"net/http"
	// Converts strings to other types and vice versa
)

func main() {
	// Use the http.NewServeMux() function to create a new servemux, which
	// is an HTTP request multiplexer. This will be used to register our
	// URL patterns and their corresponding handler functions.
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Print a message to the console indicating that the server is starting.
	log.Println("Starting server on :4000")

	// Use ListenAndServe() to start an HTTP server on port 4000, using the servemux we
	// pass in 2 params: the network address to listen on and the servemux we just created.
	// If ListenAndServe() returns an error, we log it and exit the program. Any error
	// returned from ListenAndServe is non-nil.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
