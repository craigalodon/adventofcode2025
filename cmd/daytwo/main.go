package main

import (
	"fmt"
	"math"
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

func parse(data string) ([]*mathutils.Range, error) {
	var ranges []*mathutils.Range
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
			ranges = append(ranges, mathutils.NewRange(v1, v2))
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

func intersectWithBand(ranges []*mathutils.Range, loBand, hiBand int) []*mathutils.Range {
	var out []*mathutils.Range
	for _, r := range ranges {
		lo2 := max(r.Lo, loBand)
		hi2 := min(r.Hi, hiBand)
		if lo2 <= hi2 {
			out = append(out, mathutils.NewRange(lo2, hi2))
		}
	}
	return out
}

func sumInvalids(ranges []*mathutils.Range, atLeastTwice bool) int {
	if len(ranges) == 0 {
		return 0
	}

	maxHi := 0
	for _, r := range ranges {
		if r.Hi > maxHi {
			maxHi = r.Hi
		}
	}

	if maxHi <= 0 {
		return 0
	}

	LMax := mathutils.CountDigits(maxHi)
	seen := make(map[int]bool)

	for L := 1; L < LMax+1; L++ {
		loBand := int(math.Pow(10, float64(L)-1.0))
		hiBand := int(math.Pow(10, float64(L))) - 1

		bandRanges := intersectWithBand(ranges, loBand, hiBand)

		if len(bandRanges) == 0 {
			continue
		}

		divisors := mathutils.ProperDivisors(L)

		for _, d := range divisors {
			k := L / d

			if atLeastTwice {
				if k < 2 {
					continue
				}
			} else {
				if k != 2 {
					continue
				}
			}

			R := (int(math.Pow(10, float64(L))) - 1) / (int(math.Pow(10, float64(d))) - 1)
			BDigitMin := int(math.Pow(10, float64(d)-1.0))
			BDigitMax := int(math.Pow(10, float64(d))) - 1

			for _, r := range bandRanges {
				BMin := max(BDigitMin, (r.Lo+R-1)/R)
				BMax := min(BDigitMax, r.Hi/R)

				if BMin > BMax {
					continue
				}

				for B := BMin; B <= BMax; B++ {
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
