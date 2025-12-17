package delimited

import (
	"reflect"
	"testing"
)

func TestParseSpaceDelimited(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single word",
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			name:     "two words",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "multiple spaces between words",
			input:    "hello   world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "multiple words",
			input:    "hello world test",
			expected: []string{"hello", "world", "test"},
		},
		{
			name:     "leading spaces",
			input:    "  hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "trailing spaces",
			input:    "hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "leading and trailing spaces",
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "only spaces",
			input:    "   ",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseSpaceDelimited(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseSpaceDelimited(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
