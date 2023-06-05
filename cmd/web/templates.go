package main

import (
	"github.com/xtommas/snippetbox/internal/models"
)

// holding strcuture for any dynamic data that we need to pass to the HTML templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
