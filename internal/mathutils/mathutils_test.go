package mathutils

import (
	"reflect"
	"testing"
)

func TestMatrixReduce(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name:     "Matrix 1",
			matrix:   [][]int{{0, 0, 0, 0, 1, 1}, {0, 1, 0, 0, 0, 1}, {0, 0, 1, 1, 1, 0}, {1, 1, 0, 1, 0, 0}},
			expected: [][]int{{1, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 0, 1}, {0, 0, 1, 1, 1, 0}, {0, 0, 0, 0, 1, 1}},
		},
		{
			name:     "Matrix 2",
			matrix:   [][]int{{1, 0, 1, 1, 0}, {0, 0, 0, 1, 1}, {1, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {1, 0, 1, 0, 1}},
			expected: [][]int{{1, 0, 1, 1, 0}, {0, 1, -1, 0, 1}, {0, 0, 0, 1, 1}, {0, 0, 0, 0, 2}, {0, 0, 0, 0, 0}},
		},
		{
			name:     "Matrix 3",
			matrix:   [][]int{{1, 1, 1, 0, 0, 0}, {1, 0, 1, 1, 0, 0}, {1, 1, 1, 1, 0, 0}, {1, 1, 0, 0, 0, 0}, {1, 0, 1, 0, 0, 0}, {0, 0, 1, 0, 0, 0}},
			expected: [][]int{{1, 1, 1, 0, 0, 0}, {0, -1, 0, 1, 0, 0}, {0, 0, -1, 0, 0, 0}, {0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrixReduce(tt.matrix, len(tt.matrix), len(tt.matrix[0]))
			if !reflect.DeepEqual(tt.matrix, tt.expected) {
				t.Errorf("Expected matrix %v, got %v", tt.expected, tt.matrix)
			}
		})
	}
}
