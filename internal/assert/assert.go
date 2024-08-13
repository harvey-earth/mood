package assert

import (
	"strings"
	"testing"
)

// Equal compares a real and expected test value
func Equal[T comparable](t *testing.T, real, expected T) {
	t.Helper()

	if real != expected {
		t.Errorf("test failed, real: %v; expected: %v", real, expected)
	}
}

// StringContains checks if a real string contains an expected substring value
func StringContains(t *testing.T, real, expectedSubstring string) {
	t.Helper()

	if !strings.Contains(real, expectedSubstring) {
		t.Errorf("test failed, real: %v, expected to contain: %v", real, expectedSubstring)
	}
}
