package mathutils

import (
	"errors"
	"reflect"
	"testing"
)

func TestMatrixReduce(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]float64
		expected [][]float64
		err      error
	}{
		{
			name: "Matrix 1",
			matrix: [][]float64{
				{0, 0, 0, 0, 1, 1, 3},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 1, 0, 4},
				{1, 1, 0, 1, 0, 0, 7}},
			expected: [][]float64{
				{1, 0, 0, 1, 0, -1, 2},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 0, -1, 1},
				{0, 0, 0, 0, 1, 1, 3}},
			err: nil,
		},
		{
			name: "Matrix 2",
			matrix: [][]float64{
				{1, 0, 1, 1, 0, 7},
				{0, 0, 0, 1, 1, 5},
				{1, 1, 0, 1, 1, 12},
				{1, 1, 0, 0, 1, 7},
				{1, 0, 1, 0, 1, 2}},
			expected: [][]float64{
				{1, 0, 1, 0, 0, 2},
				{0, 1, -1, 0, 0, 5},
				{0, 0, 0, 1, 0, 5},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0}},
			err: nil,
		},
		{
			name: "Matrix 3",
			matrix: [][]float64{
				{1, 1, 1, 0, 0, 0, 10},
				{1, 0, 1, 1, 0, 0, 11},
				{1, 1, 1, 1, 0, 0, 11},
				{1, 1, 0, 0, 0, 0, 5},
				{1, 0, 1, 0, 0, 0, 10},
				{0, 0, 1, 0, 0, 0, 5}},
			expected: [][]float64{
				{1, 0, 0, 0, 0, 0, 5},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 5},
				{0, 0, 0, 1, 0, 0, 1},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0}},
			err: nil,
		},
		{
			name: "Inconsistent matrix",
			matrix: [][]float64{
				{2, 10, -1},
				{3, 15, 2},
			},
			expected: nil,
			err:      errors.New("matrix is inconsistent"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := matrixReduce(tt.matrix)

			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("MatrixReduce() error = %v, expected %v", err, tt.err)
				}
			} else if err != nil {
				t.Errorf("MatrixReduce() unexpected error = %v", err)
			} else {
				if !reflect.DeepEqual(tt.matrix, tt.expected) {
					t.Errorf("MatrixReduce() = %v, want %v", tt.matrix, tt.expected)
				}
			}
		})
	}
}
