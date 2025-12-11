package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("usage: go run . <path/to/input/file> <k>")
	}

	k, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return fmt.Errorf("invalid arg <k>: %w", err)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)

	values := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		j, err := getMaxJoltage(line, k)
		if err != nil {
			return fmt.Errorf("error processing bank: %w", err)
		}
		values = append(values, j)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	sum := 0

	for _, v := range values {
		sum += v
	}

	fmt.Printf("Total joltage: %v\n", sum)
	return nil
}

func getMaxJoltage(bank string, k int) (int, error) {
	values := make([]int, k)

	for i, r := range bank {
		v, err := runeToDigit(r)
		if err != nil {
			return 0, fmt.Errorf("error parsing bank: %v", err)
		}

		x := len(bank) - i - 1

		for j, v1 := range values {
			if x < k-j-1 {
				continue
			}

			if v > v1 {
				values[j] = v
				j++
				for j < k {
					values[j] = 0
					j++
				}
				break
			}
		}
	}

	j := getJoltage(values)

	return j, nil
}

func getJoltage(values []int) int {
	i := len(values)
	i--

	j := 0
	e := 1

	for i >= 0 {
		j += values[i] * e
		e *= 10
		i--
	}

	return j
}

func runeToDigit(r rune) (int, error) {
	if unicode.IsDigit(r) {
		return int(r - '0'), nil
	}
	return 0, fmt.Errorf("rune '%c' is not a digit", r)
}
