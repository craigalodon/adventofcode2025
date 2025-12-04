package main

import (
	"fmt"
	"os"
	"strconv"

	"adventofcode2025/internal/mathutils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <path/to/input/file>")
		os.Exit(1)
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	contents := string(bytes)

	pos := 50
	stops := 0
	passes := 0
	h := 0
	t := 0

	for t < len(contents) {
		if contents[t] == '\n' {
			next, turns, err := tryMove(pos, contents[h:t])
			if err != nil {
				fmt.Printf("Conversion error: %v\n", err)
				os.Exit(1)
			}

			pos = next
			if pos == 0 {
				stops++
			}

			passes += turns

			h = t + 1
		}
		t++
	}

	fmt.Printf("Stopped on zero %v times\n", stops)
	fmt.Printf("Passed zero %v times\n", passes)
}

func tryMove(pos int, instruction string) (int, int, error) {
	n, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return pos, 0, err
	}

	turns := 0

	if instruction[0] == 'R' {
		turns += mathutils.FloorDiv(pos+n, 100)
		pos += n
	} else {
		turns += mathutils.FloorDiv(pos-n, -100) + 1
		if pos == 0 {
			turns--
		}
		pos -= n
	}

	pos = mathutils.Mod(pos, 100)

	return pos, turns, nil
}
