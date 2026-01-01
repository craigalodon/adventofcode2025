package main

import (
	"adventofcode2025/internal/mathutils"
	"adventofcode2025/internal/refutils"
	"adventofcode2025/internal/spatial"
	"bufio"
	"fmt"
	"os"
)

type Tile struct {
	X float64
	Y float64
}

func (t Tile) GetValue(depth int) float64 {

	axis := mathutils.Mod(depth, 2)

	switch axis {
	case 0:
		return t.X
	case 1:
		return t.Y
	default:
		panic("Invalid axis value")
	}
}

func (t Tile) ForEachCoordinate(fn func(dimension int, value float64)) {
	fn(0, t.X)
	fn(1, t.Y)
}

func (t Tile) SharesAxisWith(other Tile) bool {
	return t.X == other.X || t.Y == other.Y
}

func (t Tile) String() string {
	return fmt.Sprintf("(%f, %f)", t.X, t.Y)
}

func (t Tile) FormRectangle(other Tile) float64 {
	return (mathutils.Abs(t.X-other.X) + 1) * (mathutils.Abs(t.Y-other.Y) + 1)
}

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

	points, err := readPointsFromFile(os.Args[1])
	if err != nil {
		return err
	}

	pointPtrs := refutils.ToPointers(points)

	m := 0.0
	var m1 Tile
	var m2 Tile

	for i := 0; i < len(pointPtrs); i++ {
		for j := i + 1; j < len(pointPtrs); j++ {
			if pointPtrs[i].SharesAxisWith(*pointPtrs[j]) {
				continue
			}
			delta := spatial.Distance(pointPtrs[i], pointPtrs[j])
			if delta > m {
				m = delta
				m1 = *pointPtrs[i]
				m2 = *pointPtrs[j]
			}
		}
	}

	area := m1.FormRectangle(m2)

	fmt.Printf("Area: %f\n", area)

	return nil
}

func readPointsFromFile(path string) ([]Tile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	coords := make([]Tile, 0)
	for scanner.Scan() {
		var c Tile
		_, err := fmt.Sscanf(scanner.Text(), "%f,%f", &c.X, &c.Y)
		if err != nil {
			return nil, fmt.Errorf("error parsing line %q: %w", scanner.Text(), err)
		}
		coords = append(coords, c)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return coords, nil
}
