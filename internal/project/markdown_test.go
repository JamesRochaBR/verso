package project

import "testing"

func TestExtractTitle(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "simple title",
			content:  "# Hello",
			expected: "Hello",
		},
		{
			name:     "title with body",
			content:  "# My Title\n\nSome text",
			expected: "My Title",
		},
		{
			name:     "no title",
			content:  "Some text",
			expected: "",
		},
		{
			name:     "empty",
			content:  "",
			expected: "",
		},
		{
			name:     "ignore spaces",
			content:  "\n\n# Project\nBody",
			expected: "Project",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractTitle(tt.content)

			if got != tt.expected {
				t.Fatalf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
