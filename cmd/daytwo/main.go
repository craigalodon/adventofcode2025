package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"adventofcode2025/internal/mathutils"
)

type Range struct {
	Lo int
	Hi int
}

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

	ranges, err := parse(contents)

	if err != nil {
		fmt.Printf("Error parsing contents: %v\n", err)
		os.Exit(1)
	}

	part1 := sumInvalids(ranges, false)
	part2 := sumInvalids(ranges, true)

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func parse(data string) ([]Range, error) {
	ranges := []Range{}

	h2 := 0
	t2 := 0
	h1 := 0
	t1 := 0

	for t2 < len(data) {
		if data[t2] == ',' || t2+1 == len(data) {
			v1, err := strconv.Atoi(data[h1:t1])
			if err != nil {
				return nil, err
			}

			v2, err := strconv.Atoi(data[h2:t2])
			if err != nil {
				return nil, err
			}

			h2 = t2 + 1
			ranges = append(ranges, Range{v1, v2})
		}
		if data[t2] == '-' {
			h1 = h2
			t1 = t2
			h2 = t2 + 1
		}
		t2++
	}
	return ranges, nil
}

func intersectWithBand(ranges []Range, lo_band, hi_band int) []Range {
	out := []Range{}
	for _, r := range ranges {
		lo2 := max(r.Lo, lo_band)
		hi2 := min(r.Hi, hi_band)
		if lo2 <= hi2 {
			out = append(out, Range{lo2, hi2})
		}
	}
	return out
}

func sumInvalids(ranges []Range, at_least_twice bool) int {
	if len(ranges) == 0 {
		return 0
	}

	max_hi := 0
	for _, r := range ranges {
		if r.Hi > max_hi {
			max_hi = r.Hi
		}
	}

	if max_hi <= 0 {
		return 0
	}

	L_max := mathutils.CountDigits(max_hi)
	seen := make(map[int]bool)

	for L := 1; L < L_max+1; L++ {
		lo_band := int(math.Pow(10, (float64(L) - 1.0)))
		hi_band := int(math.Pow(10, float64(L))) - 1

		band_ranges := intersectWithBand(ranges, lo_band, hi_band)

		if len(band_ranges) == 0 {
			continue
		}

		divisors := mathutils.ProperDivisors(L)

		for _, d := range divisors {
			k := L / d

			if at_least_twice {
				if k < 2 {
					continue
				}
			} else {
				if k != 2 {
					continue
				}
			}

			R := (int(math.Pow(10, float64(L))) - 1) / (int(math.Pow(10, float64(d))) - 1)
			B_digit_min := int(math.Pow(10, (float64(d) - 1.0)))
			B_digit_max := int(math.Pow(10, (float64(d)))) - 1

			for _, r := range band_ranges {
				B_min := max(B_digit_min, (r.Lo+R-1)/R)
				B_max := min(B_digit_max, r.Hi/R)

				if B_min > B_max {
					continue
				}

				for B := B_min; B <= B_max; B++ {
					N := B * R
					seen[N] = true
				}
			}
		}
	}

	res := 0
	for k := range seen {
		res += k
	}

	return res
}
