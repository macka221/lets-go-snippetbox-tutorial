package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Define an application struct to hold the dependencies for the entire application.
// This will allow us to easily pass the dependencies to any part of the application.
type application struct {
	logger *slog.Logger
}

func main() {
	// Define a new CLI arg with the default value and helper text to say what
	// it does.
	addr := flag.String("addr", ":4000", "HTTP network address")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Instantiate an instance of the application struct, containing the handler dependencies.
	// In this case, we only have one dependency, which is the logger.
	app := &application{
		logger: logger,
	}
	// Before you can process the command line flags, you must first parse
	// them before using them.
	flag.Parse()

	// Use the http.NewServeMux() function to create a new servemux, which
	// is an HTTP request multiplexer. This will be used to register our
	// URL patterns and their corresponding handler functions.
	// mux := http.NewServeMux()
	//
	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	//
	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//
	// mux.HandleFunc("GET /{$}", app.home)
	// mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	// mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	// mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	//
	// Print a message to the console indicating that the server is starting.
	// log.Println("Starting server on :4000")

	// Use ListenAndServe() to start an HTTP server on port 4000, using the servemux we
	// pass in 2 params: the network address to listen on and the servemux we just created.
	// If ListenAndServe() returns an error, we log it and exit the program. Any error
	// returned from ListenAndServe is non-nil.
	// err := http.ListenAndServe(":4000", mux)

	// The value returned is a pointer to the flag value from flag.String(). We must
	// dereference it to get the actual string value.
	// log.Printf("Starting server on %s", *addr)

	logger.Info("Starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())

	// log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1)
}
