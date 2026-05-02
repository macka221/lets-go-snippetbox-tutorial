package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := httpNewServeMux()

	fileServer := httpFileServer(httpDir("./ui/static/"))
	mux.Handle("GET /static/", httpStripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
