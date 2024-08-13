package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harvey-earth/mood/internal/assert"
)

func TestSecureHeaders(t *testing.T) {
	// Initialize ResponseRecorder and http request
	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create handler to pass secureHeaders middleware
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Pass to secureHandlers and get result
	secureHeaders(next).ServeHTTP(rr, r)
	rs := rr.Result()

	// Check CSP header
	ev := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, rs.Header.Get("Content-Security-Policy"), ev)

	// Check referrer policy
	ev = "origin-when-cross-origin"
	assert.Equal(t, rs.Header.Get("Referrer-Policy"), ev)

	// Check X-Content-Type-Options
	ev = "nosniff"
	assert.Equal(t, rs.Header.Get("X-Content-Type-Options"), ev)

	// Check X-Frame-Options
	ev = "deny"
	assert.Equal(t, rs.Header.Get("X-Frame-Options"), ev)

	// Check X-XSS-Protection
	ev = "0"
	assert.Equal(t, rs.Header.Get("X-XSS-Protection"), ev)

	// Lastly check next handler in line was called and response code and body are expected
	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
