package main

import (
	"adventofcode2025/internal/mathutils"
	"bufio"
	"fmt"
	"os"
	"slices"
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

	ranges := make([]*mathutils.Range, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		r, err := parseRange(line)
		if err != nil {
			return fmt.Errorf("error parsing line: %w", err)
		}
		ranges = append(ranges, r)
	}

	slices.SortFunc(ranges, func(r1, r2 *mathutils.Range) int {
		if r1.Lo < r2.Lo {
			return -1
		}
		if r1.Lo == r2.Lo {
			if r1.Hi < r2.Hi {
				return -1
			}
			if r1.Hi == r2.Hi {
				return 0
			}
			return 1
		}
		return 1
	})

	unioned := make([]*mathutils.Range, 0, len(ranges))
	ids := 0

	curr := ranges[0]
	i := 1

	for i < len(ranges) {
		result, success := curr.Union(ranges[i])
		if success {
			curr = result
		} else {
			unioned = append(unioned, curr)
			ids += curr.Size()
			curr = ranges[i]
		}
		i++
	}

	unioned = append(unioned, curr)
	ids += curr.Size()

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		v, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("error parsing line: %w", err)
		}
		for _, r := range unioned {
			if r.Contains(v) {
				count++
				break
			}
		}
	}

	fmt.Printf("Fresh Ingredients: %v\n", count)
	fmt.Printf("There are %v ids\n", ids)

	return nil
}

func parseRange(line string) (*mathutils.Range, error) {
	for i, r := range line {
		if r == '-' {
			lo, err := strconv.Atoi(line[:i])
			if err != nil {
				return nil, fmt.Errorf("error parsing lo value: %w", err)
			}
			hi, err := strconv.Atoi(line[i+1:])
			if err != nil {
				return nil, fmt.Errorf("error parsing hi value: %w", err)
			}
			return mathutils.NewRange(lo, hi+1), nil
		}
	}
	return nil, fmt.Errorf("line does not contain a range value: %s", line)
}
