package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/xtommas/snippetbox/internal/models"
)

// the (app *application) signature means the functions are defined as a method of the application struct (sort of object oriented)
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// use the helpers
		app.NotFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// get the templateData struct containing the default data (in this case, the current year) and add the snippets slice to it
	data := app.newTemplateData(r)
	data.Snippets = snippets

	// use the render helper
	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// use the helpers
		app.NotFound(w)
		return
	}

	// Use the SnippetModel Get method to retrieve data for the record by its id, if no matching records are found, return 404 Not Found
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// get the templateData struct containing the default data (in this case, the current year) and add the snippets slice to it
	data := app.newTemplateData(r)
	data.Snippet = snippet

	// use the render helper
	app.render(w, http.StatusOK, "view.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// use the helper
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// variables holding placeholder data. Remove later
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	// pass the data to the SnippetModel.Insert() method
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
