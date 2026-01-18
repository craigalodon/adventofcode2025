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
			name: "Matrix 1",
			matrix: [][]int{
				{0, 0, 0, 0, 1, 1, 3},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 1, 0, 4},
				{1, 1, 0, 1, 0, 0, 7}},
			expected: [][]int{
				{1, 1, 0, 1, 0, 0, 7},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 1, 0, 4},
				{0, 0, 0, 0, 1, 1, 3}},
		},
		{
			name: "Matrix 2",
			matrix: [][]int{
				{1, 0, 1, 1, 0, 7},
				{0, 0, 0, 1, 1, 5},
				{1, 1, 0, 1, 1, 12},
				{1, 1, 0, 0, 1, 7},
				{1, 0, 1, 0, 1, 2}},
			expected: [][]int{
				{1, 0, 1, 1, 0, 7},
				{0, 1, -1, 0, 1, 5},
				{0, 0, 0, 1, 1, 5},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0}},
		},
		{
			name: "Matrix 3",
			matrix: [][]int{
				{1, 1, 1, 0, 0, 0, 10},
				{1, 0, 1, 1, 0, 0, 11},
				{1, 1, 1, 1, 0, 0, 11},
				{1, 1, 0, 0, 0, 0, 5},
				{1, 0, 1, 0, 0, 0, 10},
				{0, 0, 1, 0, 0, 0, 5}},
			expected: [][]int{
				{1, 1, 1, 0, 0, 0, 10},
				{0, 1, 0, -1, 0, 0, -1},
				{0, 0, 1, 0, 0, 0, 5},
				{0, 0, 0, 1, 0, 0, 1},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrixReduce(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.expected) {
				t.Errorf("Expected matrix %v, got %v", tt.expected, tt.matrix)
			}
		})
	}
}
