package mocks

import (
	"time"

	"github.com/xtommas/snippetbox/internal/models"
)

var mockSnippet = &models.Snippet{
	Id:      1,
	Title:   "No Surprises",
	Content: "A heart that's full up like a landfill, a job that slowly kills you, bruises that won't heal",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
