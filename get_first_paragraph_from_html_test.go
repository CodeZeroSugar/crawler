package main

import (
	"testing"
)

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "Get main paragraph",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
				</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "Get main paragraph multiple paragraphs",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>This is the first main paragraph.</p>
					<p>This is the second main paragraph.</p>
					<p>This is the third main paragraph.</p>
					<p>This is the fourth main paragraph.</p>
				</main>
				</body></html>`,
			expected: "This is the first main paragraph.",
		},
		{
			name: "No paragraphs",
			inputBody: `<html><body>
				<main>
				</main>
				</body></html>`,
			expected: "",
		},
		{
			name: "Get first paragraph no main.",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
					<p>This is the first main paragraph.</p>
					<p>This is the second main paragraph.</p>
					<p>This is the third main paragraph.</p>
					<p>This is the fourth main paragraph.</p>
				</body></html>`,
			expected: "Outside paragraph.",
		},
		{
			name: "Main but no paragraphs in main",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
				</main>
				<p>Second outside paragraph."
				</body></html>`,
			expected: "Outside paragraph.",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getFirstParagraphFromHTML(tc.inputBody)
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
