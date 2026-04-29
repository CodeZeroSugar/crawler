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
			inputURL: "https://www.boot.dev/blog/path",
			expected: "www.boot.dev/blog/path",
		},
		{
			name:     "ends with slash",
			inputURL: "https://www.boot.dev/blog/path/",
			expected: "www.boot.dev/blog/path",
		},
		{
			name:     "remove scheme http",
			inputURL: "http://www.boot.dev/blog/path",
			expected: "www.boot.dev/blog/path",
		},
		{
			name:     "remove scheme http and ends with slash",
			inputURL: "http://www.boot.dev/blog/path/",
			expected: "www.boot.dev/blog/path",
		},
		{
			name:     "no url",
			inputURL: "",
			expected: "",
		},
		{
			name:     "no path",
			inputURL: "http://www.boot.dev",
			expected: "www.boot.dev",
		},
		{
			name:     "no path trailing slash",
			inputURL: "http://www.boot.dev/",
			expected: "www.boot.dev",
		},
		{
			name:     "no scheme",
			inputURL: "www.boot.dev/blog/path",
			expected: "www.boot.dev/blog/path",
		},
		{
			name:     "random capitalization",
			inputURL: "http://wWW.BoOT.DEv/BLOG/pAtH",
			expected: "www.boot.dev/blog/path",
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
