package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// the (app *application) signature means the functions are defined as a method of the application struct (sort of object oriented)
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// use the helpers
		app.NotFound(w)
		return
	}

	// slice (like an array, but dynamic) of strings for the file names
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	// "files..." unpacks the elements of the slice as individual arguments to the ParseFiles function
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// use the helpers
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// use the helpers
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// use the helpers
		app.NotFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// use the helper
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
