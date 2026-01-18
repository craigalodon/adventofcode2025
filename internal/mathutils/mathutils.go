package mathutils

import (
	"fmt"
	"math"
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
	m := len(matrix)
	n := len(matrix[0])

	h := 0
	k := 0

	for h < m && k < n {
		iMax := h
		for i := h + 1; i < m; i++ {
			if math.Abs(matrix[i][k]) > math.Abs(matrix[iMax][k]) {
				iMax = i
			}
		}

		if isZero(matrix[iMax][k]) {
			k++
		} else {
			matrix[h], matrix[iMax] = matrix[iMax], matrix[h]

			for i := h + 1; i < m; i++ {
				f := matrix[i][k] / matrix[h][k]
				matrix[i][k] = 0
				for j := k + 1; j < n; j++ {
					matrix[i][j] = matrix[i][j] - matrix[h][j]*f
				}
			}

			for i := n - 1; i >= k; i-- {
				matrix[h][i] = matrix[h][i] / matrix[h][k]
			}

			h++
			k++
		}
	}

	h--

	for h > 0 {
		k = 0
		for range n {
			if matrix[h][k] == 1 {
				break
			}
			k++
		}

		if isZero(matrix[h][k]) {
			h--
			continue
		}

		if k == n-1 {
			return fmt.Errorf("matrix is inconsistent")
		}

		for i := 0; i < h; i++ {
			if !isZero(matrix[i][k]) {
				f := matrix[i][k]
				for j := k; j < n; j++ {
					matrix[i][j] = matrix[i][j] - matrix[h][j]*f
				}
			}
		}

		h--
	}

	return nil
}

func isZero(f float64) bool {
	const epsilon = 1e-10
	return math.Abs(f) < epsilon
}
