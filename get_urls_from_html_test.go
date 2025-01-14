package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "only absolute URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="https://example.com/path/one">
			<span>Example</span>
		</a>
		<a href="https://example.com/path/two">
			<span>Example</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://example.com/path/one", "https://example.com/path/two"},
		},
		{
			name:     "only relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="/path/two">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
		},
		{
			name:     "no URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<p>No links here!</p>
	</body>
</html>
`,
			expected: []string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test '%s' FAIL: unexpected error: %v", tc.name, err)
				return
			}
			if len(tc.expected) == 0 && len(actual) == 0 {
				return // both are empty slices, consider them equal
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test '%s' FAIL: expected URLs: %v, actual: %v", tc.name, tc.expected, actual)
			}
		})
	}
}
