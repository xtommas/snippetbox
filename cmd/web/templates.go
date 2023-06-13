package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/xtommas/snippetbox/internal/models"
)

// holding strcuture for any dynamic data that we need to pass to the HTML templates
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
}

// returns a nicely formatted string representation of the time
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// initialize the template.FuncMap object and store in a global variable
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// get a slice of all filepaths that match the pattern "./ui/html/pages/*.html"
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// loop through the filepaths
	for _, page := range pages {
		// extract the name from the full filepath
		name := filepath.Base(page)

		// parse the files into a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// call ParseGlob() to add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// add the page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// add the template set to the map, using the naeme of the page as the key
		cache[name] = ts
	}

	return cache, nil
}
