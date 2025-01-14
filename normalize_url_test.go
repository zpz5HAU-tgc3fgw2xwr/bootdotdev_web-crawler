package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove www",
			inputURL: "https://www.boot.dev/path",
			expected: "boot.dev/path",
		},
		{
			name:     "http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "no scheme",
			inputURL: "blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "root path",
			inputURL: "https://www.boot.dev/",
			expected: "boot.dev/",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
