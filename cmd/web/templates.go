package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/xtommas/snippetbox/internal/models"
	"github.com/xtommas/snippetbox/ui"
)

// holding strcuture for any dynamic data that we need to pass to the HTML templates
type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

// returns a nicely formatted string representation of the time
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

// initialize the template.FuncMap object and store in a global variable
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// get a slice of all filepaths that match the pattern "./ui/html/pages/*.html" from the embedded filesystem
	pages, err := fs.Glob(ui.Files, "./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// loop through the filepaths
	for _, page := range pages {
		// extract the name from the full filepath
		name := filepath.Base(page)

		// slice containing the filepath patterns for the templates we want to parse
		patterns := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		// parse the template files from the embedded filesystem
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		// add the template set to the map, using the naeme of the page as the key
		cache[name] = ts
	}

	return cache, nil
}
