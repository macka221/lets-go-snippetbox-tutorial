package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// the string "Hello from Snippetbox" as the response body.
//
// (app *application): Makes this method defined against *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Adds a header to the response.
	w.Header().Add("Server", "Go")
	// w.Write([]byte("Hello from Snippetbox"))

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If theres an errror, we log the detailed error message and
	// send a generic 500 Internal Server with http.Error()
	//
	// NOTE: paths are either absolute or relative to pwd
	// ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")

	// Initialize a slice of strings containing the paths to the template files.
	// The file containing the base template must be first.
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Use the template.ParseFiles() function to read the template files and
	// store the templates in a template set, we use the '...' to pass the
	// contents of the files slice as a variadic, or unknown number of arguments.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// log.Print(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		//
		// Because the home handler is now a method against the application  struct
		// it can access its fields, including the structured logger. We'll use this
		// to create a log entry at Error level container the error message, also
		// including the request method and URI as attributes to assist w/ debugging.
		// app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err)
		return
	}

	// Then we use the Execute() method on the template set to write the
	// template content as a response body.
	// NOTE: The last parameter of  Execute() is the dynamic data that
	// we want to pass to the template.
	// err = ts.Execute(w, nil)

	// Use the ExecuteTemplate method to write the content of the base template as the
	// response body. This method takes an additional parameter which is the name of
	// the template we want to execute.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// log.Print(err.Error())
		// app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err)
	}
}

// Add a snippertView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
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
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Customizes the headers returned by this method
	w.WriteHeader(http.StatusCreated) // http has constants that represents all status codes.
	w.Write([]byte("Save a new snippet..."))
}
