package main

import (
	"errors"
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
				indicators:          []bool{true, false, false, true},
				buttons:             []map[int]bool{{1: true, 3: true}, {2: true}, {2: true, 3: true}, {0: true, 2: true}, {0: true, 1: true}, {3: true}},
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
