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

func MatrixReduce(matrix [][]float64) ([][]float64, error) {

	rref := make([][]float64, len(matrix))
	for i := range rref {
		rref[i] = make([]float64, len(matrix[i]))
		copy(rref[i], matrix[i])
	}

	h := 0
	k := 0

	for h < len(rref) && k < len(rref[0]) {
		iMax := h
		for i := h + 1; i < len(rref); i++ {
			if math.Abs(rref[i][k]) > math.Abs(rref[iMax][k]) {
				iMax = i
			}
		}

		if IsZero(rref[iMax][k]) {
			k++
		} else {
			rref[h], rref[iMax] = rref[iMax], rref[h]

			for i := h + 1; i < len(rref); i++ {
				f := rref[i][k] / rref[h][k]
				rref[i][k] = 0
				for j := k + 1; j < len(rref[0]); j++ {
					rref[i][j] = rref[i][j] - rref[h][j]*f
				}
			}

			for i := len(rref[0]) - 1; i >= k; i-- {
				rref[h][i] = rref[h][i] / rref[h][k]
			}

			h++
			k++
		}
	}

	h--

	for h > 0 {
		k = 0
		for range len(rref[0]) {
			if rref[h][k] == 1 {
				break
			}
			k++
		}

		if IsZero(rref[h][k]) {
			h--
			continue
		}

		if k == len(rref[0])-1 {
			return nil, fmt.Errorf("matrix is inconsistent")
		}

		for i := 0; i < h; i++ {
			if !IsZero(rref[i][k]) {
				f := rref[i][k]
				for j := k; j < len(rref[0]); j++ {
					rref[i][j] = rref[i][j] - rref[h][j]*f
				}
			}
		}

		h--
	}

	return rref, nil
}

func findPivots(matrix [][]float64) map[int]int {
	pivots := make(map[int]int)

	h := 0
	k := 0

	for h < len(matrix) && k < len(matrix[0])-1 {
		if IsZero(matrix[h][k]) {
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
			if !IsZero(matrix[h][j]) {
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

func Parametrize(matrix [][]float64) (map[int]int, [][]float64) {
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

	return freeVariables, params
}

func IsZero(f float64) bool {
	const epsilon = 1e-10
	return math.Abs(f) < epsilon
}

func GetCoefficients(params [][]float64, args []float64) []float64 {
	coefs := make([]float64, len(params))
	for i, exp := range params {
		acc := 0.0
		for j, arg := range args {
			acc += exp[j] * arg
		}
		coefs[i] = acc
	}
	return coefs
}

func CoefsConsistentWithMatrix(matrix [][]float64, coefs []float64) bool {
	valid := true
	for _, row := range matrix {
		acc := 0.0
		for i, coef := range coefs {
			acc += row[i] * coef
		}
		if !IsZero(acc - row[len(row)-1]) {
			valid = false
			break
		}
	}
	return valid
}
