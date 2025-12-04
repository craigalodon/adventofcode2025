package mathutils

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
	res := []int{}

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
