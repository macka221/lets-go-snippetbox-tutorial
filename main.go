package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv" // Converts strings to other types and vice versa
)

// Define a home handler function which writes a byte slice containing
// the string "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Adds a header to the response.
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippertView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	// Validates the value of the error and id. NOTE: Go does not have exceptions.
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
	// msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	// w.Write([]byte(msg))
}

// Add a snippertCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Customizes the headers returned by this method
	w.WriteHeader(http.StatusCreated) // http has constants that represents all status codes.
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to create a new servemux, which
	// is an HTTP request multiplexer. This will be used to register our
	// URL patterns and their corresponding handler functions.
	mux := http.NewServeMux()
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
