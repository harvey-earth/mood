package main

import (
	"net/http"
	"testing"

	"github.com/harvey-earth/mood/internal/assert"
)

func TestPing(t *testing.T) {
	// Instantiate app
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Make GET request to /ping endpoint
	code, _, body := ts.get(t, "/ping")

	// Test results
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "pong")
}

func TestTeamView(t *testing.T) {
	// Instantiate app
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Setup testing struct for team view
	tests := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Good ID",
			path:         "/teams/1",
			expectedCode: http.StatusOK,
			expectedBody: "Example Team",
		},
		{
			name:         "Bad ID",
			path:         "/teams/2",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "String ID",
			path:         "/teams/foo",
			expectedCode: http.StatusNotFound,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.path)

			assert.Equal(t, code, tt.expectedCode)

			if tt.expectedBody != "" {
				assert.StringContains(t, body, tt.expectedBody)
			}
		})
	}
}
