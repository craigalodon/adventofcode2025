package mathutils

import "math"

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

func matrixReduce(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	h := 0
	k := 0

	for h < m && k < n {
		iMax := h
		for i := h + 1; i < m; i++ {
			if math.Abs(float64(matrix[i][k])) > math.Abs(float64(matrix[iMax][k])) {
				iMax = i
			}
		}

		if matrix[iMax][k] == 0 {
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

		if matrix[h][k] == 0 {
			h--
			continue
		}

		for i := 0; i < h; i++ {
			if matrix[i][k] != 0 {
				f := matrix[i][k]
				for j := 0; j < n; j++ {
					matrix[i][j] = matrix[i][j] - matrix[h][j]*f
				}
			}
		}

		h--
	}
}
