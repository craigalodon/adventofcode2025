package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *Machine
		err      error
	}{
		{
			name:  "valid machine input",
			input: "[#..#] (1,3) (2) (2,3) (0,2) (0,1) (3) {7,4,3,5}",
			expected: &Machine{
				configuration:       9,
				buttons:             []int{12, 10, 8, 5, 4, 3},
				joltageRequirements: map[int]bool{7: true, 4: true, 3: true, 5: true},
			},
			err: nil,
		},
		{
			name:     "unexpected character at the beginning",
			input:    " [#..#] (1,3) (2) (2,3) (0,2) (0,1) (3) {7,4,3,5}",
			expected: nil,
			err:      errors.New("invalid character ' ' at position 0"),
		},
		{
			name:     "unexpected character at the end",
			input:    "[#..#] (1,3) (2) (2,3) (0,2) (0,1) (3) {7,4,3,5}\n",
			expected: nil,
			err:      errors.New("invalid character at position 48"),
		},
		{
			name:     "unexpected character in brackets",
			input:    "[#,..#] (1,3) (2) (2,3) (0,2) (0,1) (3) {7,4,3,5}",
			expected: nil,
			err:      errors.New("invalid character ',' at position 2"),
		},
		{
			name:     "repeated commas",
			input:    "[#..#] (1,,3) (2) (2,3) (0,2) (0,1) (3) {7,4,3,5}",
			expected: nil,
			err:      errors.New("strconv.Atoi: parsing \"\": invalid syntax"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := deserialize(tt.input)

			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("deserialize() error = %v, expected %v", err, tt.err)
				}
			} else if err != nil {
				t.Errorf("deserialize() unexpected error = %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("deserialize() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAddIndicator(t *testing.T) {
	tests := []struct {
		name      string
		button    int
		indicator int
		expected  int
	}{
		{
			name:      "button 0, indicator 1",
			button:    0,
			indicator: 1,
			expected:  2,
		},
		{
			name:      "button 0, indicator 2",
			button:    0,
			indicator: 2,
			expected:  4,
		},
		{
			name:      "button 1, indicator 1",
			button:    1,
			indicator: 1,
			expected:  3,
		},
		{
			name:      "button 1, indicator 2",
			button:    1,
			indicator: 2,
			expected:  5,
		},
		{
			name:      "button 2, indicator 1",
			button:    2,
			indicator: 1,
			expected:  2,
		},
		{
			name:      "button 2, indicator 2",
			button:    2,
			indicator: 2,
			expected:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addIndicator(tt.button, tt.indicator)

			if result != tt.expected {
				t.Errorf("addIndicator() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestConfigure(t *testing.T) {
	tests := []struct {
		name       string
		serialized string
		expected   int
	}{
		{
			name:     "Example 1",
			expected: 2,
		},
		{
			name:     "Example 2",
			expected: 3,
		},
		{
			name:     "Example 3",
			expected: 2,
		},
	}

	filepath := "../../testdata/dayten/example.txt"
	file, err := os.Open(filepath)
	if err != nil {
		t.Fatalf("error opening file: %v", filepath)
	}
	defer func() {
		if err := file.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !scanner.Scan() {
				t.Fatalf("unexpected end of file")
			}
			machine, err := deserialize(scanner.Text())
			if err != nil {
				t.Fatalf("deserialize() error = %v", err)
			}
			result := machine.configure()
			if result != tt.expected {
				t.Errorf("configure() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPress(t *testing.T) {
	tests := []struct {
		name     string
		curr     int
		button   int
		expected int
	}{
		{
			name:     "curr 0, button 0",
			curr:     0,
			button:   0,
			expected: 0,
		},
		{
			name:     "curr 0, button 1",
			curr:     0,
			button:   1,
			expected: 1,
		},
		{
			name:     "curr 1, button 1",
			curr:     1,
			button:   1,
			expected: 0,
		},
		{
			name:     "curr 1, button 2",
			curr:     1,
			button:   2,
			expected: 3,
		},
		{
			name:     "curr 2, button 2",
			curr:     2,
			button:   2,
			expected: 0,
		},
		{
			name:     "curr 2, button 3",
			curr:     2,
			button:   3,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := press(tt.curr, tt.button)
			if result != tt.expected {
				t.Errorf("press() = %v, want %v", result, tt.expected)
			}
		})
	}
}
