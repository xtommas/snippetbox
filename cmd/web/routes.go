package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// returns a servemux containing the routes of the app
func (app *application) routes() http.Handler {
	router := httprouter.New()

	// create a handler function that wraps notFound() helper, then assign it as the custom handler for 404 responses
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.NotFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	// create a middleware chain
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// pass the servemux as the 'next' parameter to the middleware
	return standard.Then(router)
}
