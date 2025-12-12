package main

import (
	"fmt"
	"os"

	"adventofcode2025/internal/mathutils"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <path/to/input/file> <retries>")
		os.Exit(1)
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	retries := os.Args[2] == "true"

	contents := string(bytes)

	cols := getCols(contents)
	seen, rolls := getRolls(contents, cols)

	movable := 0
	moved := make(map[int]bool)
	try := true

	for try {
		try = false
		for _, r := range rolls {
			if moved[r] {
				continue
			}

			adj := getAdjacent(r, cols, seen)

			if adj < 4 {
				movable++
				moved[r] = true
				try = true
			}
		}

		if !retries {
			break
		}

		// remove moved
		for k := range moved {
			seen[k] = false
		}
	}

	fmt.Printf("Movable Rolls: %v\n", movable)
}

func getAdjacent(r int, cols int, seen map[int]bool) int {
	adj := 0

	e := r + 1
	w := r - 1
	n := r - cols
	s := r + cols
	ne := n + 1
	nw := n - 1
	se := s + 1
	sw := s - 1

	left := mathutils.Mod(r, cols) != 0
	right := mathutils.Mod(r+1, cols) != 0

	if right && seen[e] {
		adj++
	}
	if left && seen[w] {
		adj++
	}
	if seen[n] {
		adj++
	}
	if seen[s] {
		adj++
	}
	if right && seen[ne] {
		adj++
	}
	if left && seen[nw] {
		adj++
	}
	if right && seen[se] {
		adj++
	}
	if left && seen[sw] {
		adj++
	}
	return adj
}

func getRolls(contents string, cols int) (map[int]bool, []int) {
	seen := make(map[int]bool)
	rolls := make([]int, 0)

	i := 0
	j := 0

	for _, r := range contents {
		if r == '\n' {
			i++
			j = 0
			continue
		}

		if r == '@' {
			seen[i*cols+j] = true
			rolls = append(rolls, i*cols+j)
		}

		j++
	}
	return seen, rolls
}

func getCols(contents string) int {
	cols := 0

	for _, r := range contents {
		if r == '\n' {
			break
		}
		cols++
	}
	return cols
}
