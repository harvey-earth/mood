package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harvey-earth/mood/internal/models/mocks"
)

// Helper to create new test application instance
func newTestApplication(t *testing.T) *application {
	return &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
		teams:    &mocks.TeamModel{},
	}
}

// Struct to hold testServer
type testServer struct {
	*httptest.Server
}

// Helper to create new test server
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

// Helper to run a GET request against an endpoint
// Returns response status code, header, and body
func (ts *testServer) get(t *testing.T, url string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + url)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}
