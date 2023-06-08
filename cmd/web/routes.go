package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// returns a servemux containing the routes of the app
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// create a middleware chain
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// pass the servemux as the 'next' parameter to the middleware
	return standard.Then(mux)
}
