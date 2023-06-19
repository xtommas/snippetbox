package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet model
type Snippet struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Interfaces define the "methods" that the Snippet struct has, kinda like OOP. Here it's useful to allow mocking dependencies in the application struct
type SnippetModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (*Snippet, error)
	Latest() ([]*Snippet, error)
}

// snippetModel type that wraps the sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert a snippet in the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// use placeholder parameters (?) to avoid SQL injection
	query := "INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))"

	// Execute the query with the exec method (used for queries that don't return rows)
	result, err := m.DB.Exec(query, title, content, expires)
	if err != nil {
		return 0, err
	}

	// get the id of the inserted record to return it
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// convert the id from int4 to int
	return int(id), nil
}

// Return a snippet by id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	query := "SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?"

	// use the QueryRow for queries that return a single row
	row := m.DB.QueryRow(query, id)

	// initialize a pointer to an empty Snippet struct
	s := &Snippet{}

	// copy the values from the row to the Snippet object
	err := row.Scan(&s.Id, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// if the query returns no rows, row.Scan() returns sql.ErrNoRows error
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// return the Snippet object
	return s, nil
}

// Return the last 10 snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	query := "SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10"

	// use the Query() method for mulitple row responses from db
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	// defer to ensure the sql.Rows resultset is properly closed before the Latest() method returns to avoid using all the database pool connections
	defer rows.Close()

	// empty slice of emtpy Snippet structs
	snippets := []*Snippet{}

	// use rows.Next to iterate through the rows in the resultset
	for rows.Next() {
		// initialize emtpy snippet struct object
		s := &Snippet{}

		// map it to the row response
		err = rows.Scan(&s.Id, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		// add it to the snippets slice
		snippets = append(snippets, s)
	}

	// check for errors when the rows.Next() loop finishes
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// return the snippets slice
	return snippets, nil
}
