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
	rolls := getRolls(contents, cols)

	movable := 0
	try := true

	for try {
		try = false
		removed := make([]int, 0)

		for r, exists := range rolls {
			if !exists {
				continue
			}

			adj := getAdjacent(r, cols, rolls)

			if adj < 4 {
				movable++
				removed = append(removed, r)
				try = true
			}
		}

		if !retries {
			break
		}

		for _, r := range removed {
			rolls[r] = false
		}
	}

	fmt.Printf("Movable Rolls: %v\n", movable)
}

func getAdjacent(r int, cols int, rolls map[int]bool) int {
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

	if right && rolls[e] {
		adj++
	}
	if left && rolls[w] {
		adj++
	}
	if rolls[n] {
		adj++
	}
	if rolls[s] {
		adj++
	}
	if right && rolls[ne] {
		adj++
	}
	if left && rolls[nw] {
		adj++
	}
	if right && rolls[se] {
		adj++
	}
	if left && rolls[sw] {
		adj++
	}
	return adj
}

func getRolls(contents string, cols int) map[int]bool {
	rolls := make(map[int]bool)

	i := 0
	j := 0

	for _, r := range contents {
		if r == '\n' {
			i++
			j = 0
			continue
		}

		if r == '@' {
			rolls[i*cols+j] = true
		}

		j++
	}
	return rolls
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
