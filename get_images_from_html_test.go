package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetImagesFromHTMLRelative(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "Get one image",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "Get two images",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"><img src="/boots.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png", "https://crawler-test.com/boots.png"},
		},
		{
			name:      "No images",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><p>No Images here...</p></body></html>`,
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
			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
