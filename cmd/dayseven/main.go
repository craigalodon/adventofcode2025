package main

import (
	"bufio"
	"fmt"
	"os"
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

	beams := make(map[int]bool)
	splits := 0
	for scanner.Scan() {
		line := scanner.Text()
		next := make(map[int]bool)
		for k, v := range beams {
			next[k] = v
		}
		for i, r := range line {
			if r == 'S' {
				next[i] = true
				continue
			}
			if r == '^' && beams[i] {
				splits++
				next[i] = false
				if i-1 >= 0 {
					next[i-1] = true
				}
				if i+1 < len(line) {
					next[i+1] = true
				}
			}
		}
		beams = next
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	fmt.Printf("Splits: %v\n", splits)

	return nil
}
