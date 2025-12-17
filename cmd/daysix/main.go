package main

import (
	"adventofcode2025/internal/delimited"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: go run . <path/to/input/file>")
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

	vals := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := delimited.ParseSpaceDelimited(line)
		ints, success := tryConvertRowToInts(row)
		if success {
			vals = append(vals, ints)
		} else {
			sum, err := aggregate(row, vals)
			if err != nil {
				return err
			}
			fmt.Printf("Sum: %v\n", sum)
		}
	}

	return nil
}

func tryConvertRowToInts(row []string) ([]int, bool) {
	vals := make([]int, 0)
	for _, s := range row {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, false
		} else {
			vals = append(vals, v)
		}
	}

	return vals, true
}

func aggregate(ops []string, vals [][]int) (int, error) {
	sum := 0

	for i, r := range ops {
		if r == "+" {
			acc := 0
			for _, row := range vals {
				acc += row[i]
			}
			sum += acc
			continue
		}

		if r == "*" {
			acc := 1
			for _, row := range vals {
				acc *= row[i]
			}
			sum += acc
			continue
		}

		return 0, fmt.Errorf("unexpected operator: %s", r)
	}

	return sum, nil
}
