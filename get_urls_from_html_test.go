package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "Get Url",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com"},
		},
		{
			name:      "Get Urls",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com", "https://crawler-test.com"},
		},
		{
			name:      "Get Urls, with a relative",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a><a href="/posts/article"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com", "https://crawler-test.com/posts/article"},
		},
		{
			name:      "No URLs",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><p>No URLs here</p></body></html>`,
			expected:  []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
