package mathutils

import (
	"fmt"
	"math"
	"slices"
)

func Mod(a, b int) int {
	return ((a % b) + b) % b
}

func FloorDiv(a, b int) int {
	if a%b == 0 {
		return a / b
	}

	if (a < 0) != (b < 0) {
		return a/b - 1
	}
	return a / b
}

func ProperDivisors(n int) []int {
	var res []int

	i := 1

	for i*i <= n {
		if n%i == 0 {
			if i < n {
				res = append(res, i)
			}
			j := n / i
			if j != i && j < n {
				res = append(res, j)
			}
		}
		i++
	}

	return res
}

func CountDigits(n int) int {
	if n < 0 {
		n = -n
	}

	if n == 0 {
		return 1
	}

	d := 0
	for n > 0 {
		n /= 10
		d++
	}

	return d
}

func matrixReduce(matrix [][]float64) error {
	h := 0
	k := 0

	for h < len(matrix) && k < len(matrix[0]) {
		iMax := h
		for i := h + 1; i < len(matrix); i++ {
			if math.Abs(matrix[i][k]) > math.Abs(matrix[iMax][k]) {
				iMax = i
			}
		}

		if isZero(matrix[iMax][k]) {
			k++
		} else {
			matrix[h], matrix[iMax] = matrix[iMax], matrix[h]

			for i := h + 1; i < len(matrix); i++ {
				f := matrix[i][k] / matrix[h][k]
				matrix[i][k] = 0
				for j := k + 1; j < len(matrix[0]); j++ {
					matrix[i][j] = matrix[i][j] - matrix[h][j]*f
				}
			}

			for i := len(matrix[0]) - 1; i >= k; i-- {
				matrix[h][i] = matrix[h][i] / matrix[h][k]
			}

			h++
			k++
		}
	}

	h--

	for h > 0 {
		k = 0
		for range len(matrix[0]) {
			if matrix[h][k] == 1 {
				break
			}
			k++
		}

		if isZero(matrix[h][k]) {
			h--
			continue
		}

		if k == len(matrix[0])-1 {
			return fmt.Errorf("matrix is inconsistent")
		}

		for i := 0; i < h; i++ {
			if !isZero(matrix[i][k]) {
				f := matrix[i][k]
				for j := k; j < len(matrix[0]); j++ {
					matrix[i][j] = matrix[i][j] - matrix[h][j]*f
				}
			}
		}

		h--
	}

	return nil
}

func findPivots(matrix [][]float64) map[int]int {
	pivots := make(map[int]int)

	h := 0
	k := 0

	for h < len(matrix) && k < len(matrix[0])-1 {
		if isZero(matrix[h][k]) {
			k++
		} else {
			pivots[k] = h
			h++
			k++
		}
	}
	return pivots
}

func findFreeVariables(matrix [][]float64, pivots map[int]int) map[int]int {
	freeVariables := make(map[int]int)
	arr := make([]int, 0)
	for k, h := range pivots {
		for j := k + 1; j < len(matrix[0])-1; j++ {
			if !isZero(matrix[h][j]) {
				if _, ok := freeVariables[j]; !ok {
					freeVariables[j] = 0
					arr = append(arr, j)
				}
			}
		}
	}

	slices.Sort(arr)
	for i, j := range arr {
		freeVariables[j] = i
	}

	return freeVariables
}

func parametrize(matrix [][]float64) [][]float64 {
	pivots := findPivots(matrix)
	freeVariables := findFreeVariables(matrix, pivots)

	params := make([][]float64, len(matrix[0])-1)
	for i := 0; i < len(params); i++ {
		params[i] = make([]float64, len(freeVariables)+1)
	}

	for k, h := range pivots {
		params[k][0] = matrix[h][len(matrix[0])-1]
		for i, j := range freeVariables {
			if i > k {
				params[k][j+1] = matrix[h][i] * -1
			}
		}
	}

	for i, j := range freeVariables {
		params[i][j+1] = 1
	}

	return params
}

func isZero(f float64) bool {
	const epsilon = 1e-10
	return math.Abs(f) < epsilon
}
