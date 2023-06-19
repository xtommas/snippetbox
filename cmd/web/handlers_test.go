package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/xtommas/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	// new instance of the application struct. For now, just contains mock loggers that discard anything written to them
	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}

	// create a new test server
	ts := httptest.NewTLSServer(app.routes())
	// defer a call to close the server when the test finishes
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
