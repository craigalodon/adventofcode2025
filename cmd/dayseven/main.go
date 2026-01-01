package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
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
			_, _ = fmt.Fprintf(os.Stderr, "Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)

	beams := make(map[int]int)
	splits := 0
	timelines := 1
	for scanner.Scan() {
		line := scanner.Text()
		next := make(map[int]int)
		maps.Copy(next, beams)
		for i, r := range line {
			if r == 'S' {
				next[i] = 1
				continue
			}
			if r == '^' && beams[i] > 0 {
				splits++
				timelines += beams[i]
				next[i] = 0
				if i-1 >= 0 {
					next[i-1] += beams[i]
				}
				if i+1 < len(line) {
					next[i+1] += beams[i]
				}
			}
		}
		beams = next
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	fmt.Printf("Splits: %v\n", splits)
	fmt.Printf("Timelines: %v\n", timelines)

	return nil
}
