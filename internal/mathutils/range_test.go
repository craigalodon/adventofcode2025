package mathutils

import (
	"reflect"
	"testing"
)

func TestGenerateCombinations(t *testing.T) {
	tests := []struct {
		name   string
		ranges []*Range
		expect [][]int
	}{
		{
			name: "3 x 2",
			ranges: []*Range{
				{
					Lo: 0,
					Hi: 2,
				},
				{
					Lo: 0,
					Hi: 2,
				},
				{
					Lo: 0,
					Hi: 2,
				},
			},
			expect: [][]int{
				{0, 0, 0},
				{0, 0, 1},
				{0, 1, 0},
				{0, 1, 1},
				{1, 0, 0},
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 1},
			},
		},
		{
			name: "2 x 3",
			ranges: []*Range{
				{
					Lo: 0,
					Hi: 3,
				},
				{
					Lo: 0,
					Hi: 3,
				},
			},
			expect: [][]int{
				{0, 0},
				{0, 1},
				{0, 2},
				{1, 0},
				{1, 1},
				{1, 2},
				{2, 0},
				{2, 1},
				{2, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateCombinations(tt.ranges)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("GenerateCombinations(%v) = %v, want %v", tt.ranges, result, tt.expect)
			}
		})
	}
}
