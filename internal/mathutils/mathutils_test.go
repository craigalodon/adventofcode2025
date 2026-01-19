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
				{1, 1, 1, 0, 10},
				{1, 0, 1, 1, 11},
				{1, 1, 1, 1, 11},
				{1, 1, 0, 0, 5},
				{1, 0, 1, 0, 10},
				{0, 0, 1, 0, 5}},
			expected: [][]float64{
				{1, 0, 0, 0, 5},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 0, 5},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0}},
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
		{
			name: "Matrix 4",
			matrix: [][]float64{
				{1, 1, 1, 0, 10},
				{1, 0, 1, 1, 11},
				{1, 0, 1, 1, 11},
				{1, 1, 0, 0, 5},
				{1, 1, 1, 0, 10},
				{0, 0, 1, 0, 5}},
			expected: [][]float64{
				{1, 0, 0, 1, 6},
				{0, 1, 0, -1, -1},
				{0, 0, 1, 0, 5},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0}},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rref, err := MatrixReduce(tt.matrix)

			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("MatrixReduce() error = %v, expected %v", err, tt.err)
				}
			} else if err != nil {
				t.Errorf("MatrixReduce() unexpected error = %v", err)
			} else {
				if !reflect.DeepEqual(rref, tt.expected) {
					t.Errorf("MatrixReduce() = %v, want %v", rref, tt.expected)
				}
			}
		})
	}
}

func TestFindPivots(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]float64
		expected map[int]int
	}{
		{
			name: "Matrix 1",
			matrix: [][]float64{
				{1, 0, 0, 1, 0, -1, 2},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 0, -1, 1},
				{0, 0, 0, 0, 1, 1, 3}},
			expected: map[int]int{
				0: 0,
				1: 1,
				2: 2,
				4: 3},
		},
		{
			name: "Matrix 2",
			matrix: [][]float64{
				{1, 0, 1, 0, 0, 2},
				{0, 1, -1, 0, 0, 5},
				{0, 0, 0, 1, 0, 5},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0}},
			expected: map[int]int{
				0: 0,
				1: 1,
				3: 2,
				4: 3},
		},
		{
			name: "Matrix 3",
			matrix: [][]float64{
				{1, 0, 0, 0, 5},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 0, 5},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0}},
			expected: map[int]int{
				0: 0,
				1: 1,
				2: 2,
				3: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pivots := findPivots(tt.matrix)
			if !reflect.DeepEqual(pivots, tt.expected) {
				t.Errorf("findPivots() = %v, want %v", pivots, tt.expected)
			}
		})
	}
}

func TestFindFreeVariables(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]float64
		pivots   map[int]int
		expected map[int]int
	}{
		{
			name: "Matrix 1",
			matrix: [][]float64{
				{1, 0, 0, 1, 0, -1, 2},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 0, -1, 1},
				{0, 0, 0, 0, 1, 1, 3}},
			pivots: map[int]int{
				0: 0,
				1: 1,
				2: 2,
				4: 3},
			expected: map[int]int{
				3: 0,
				5: 1},
		},
		{
			name: "Matrix 2",
			matrix: [][]float64{
				{1, 0, 1, 0, 0, 2},
				{0, 1, -1, 0, 0, 5},
				{0, 0, 0, 1, 0, 5},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0}},
			pivots: map[int]int{
				0: 0,
				1: 1,
				3: 2,
				4: 3},
			expected: map[int]int{2: 0},
		},
		{
			name: "Matrix 3",
			matrix: [][]float64{
				{1, 0, 0, 0, 5},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 0, 5},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0}},
			pivots: map[int]int{
				0: 0,
				1: 1,
				2: 2,
				3: 3},
			expected: map[int]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			freeVariables := findFreeVariables(tt.matrix, tt.pivots)
			if !reflect.DeepEqual(freeVariables, tt.expected) {
				t.Errorf("findFreeVariables() = %v, want %v", freeVariables, tt.expected)
			}
		})
	}
}

func TestParametrize(t *testing.T) {
	tests := []struct {
		name      string
		matrix    [][]float64
		variables map[int]int
		expected  [][]float64
	}{
		{
			name: "Matrix 1",
			matrix: [][]float64{
				{1, 0, 0, 1, 0, -1, 2},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 0, -1, 1},
				{0, 0, 0, 0, 1, 1, 3}},
			variables: map[int]int{
				3: 0,
				5: 1},
			expected: [][]float64{
				{2, -1, 1},
				{5, 0, -1},
				{1, -1, 1},
				{0, 1, 0},
				{3, 0, -1},
				{0, 0, 1},
			},
		},
		{
			name: "Matrix 2",
			matrix: [][]float64{
				{1, 0, 1, 0, 0, 2},
				{0, 1, -1, 0, 0, 5},
				{0, 0, 0, 1, 0, 5},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0}},
			variables: map[int]int{2: 0},
			expected: [][]float64{
				{2, -1},
				{5, 1},
				{0, 1},
				{5, 0},
				{0, 0},
			},
		},
		{
			name: "Matrix 3",
			matrix: [][]float64{
				{1, 0, 0, 0, 5},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 0, 5},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0}},
			variables: map[int]int{},
			expected: [][]float64{
				{5},
				{0},
				{5},
				{1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, params := Parametrize(tt.matrix)
			if !reflect.DeepEqual(params, tt.expected) {
				t.Errorf("Parametrize() = %v, want %v", params, tt.expected)
			}
		})
	}
}
