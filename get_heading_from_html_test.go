package main

import (
	"testing"
)

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "basic",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "basic h2",
			inputBody: "<html><body><h2>Backup Title</h2></body></html>",
			expected:  "Backup Title",
		},
		{
			name:      "no h1 or h2",
			inputBody: "<html><body><main><p>There is some text here to distract you.</p></main></body></html>",
			expected:  "",
		},
		{
			name:      "h1 blank with h2 text",
			inputBody: "<html><body><h1></h1><main><p>There is some text here to distract you.</p></main><h2>Here is what you want</h2><p>Another distraction</p></body></html>",
			expected:  "Here is what you want",
		},
		{
			name:      "invisible header",
			inputBody: "<html><body><h2>                                                   </h2></body></html>",
			expected:  "",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getHeadingFromHTML(tc.inputBody)
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
