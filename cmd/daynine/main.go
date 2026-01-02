package main

import (
	"adventofcode2025/internal/refutils"
	"bufio"
	"fmt"
	"math"
	"os"
)

type Tile struct {
	X float64
	Y float64
}

func (t Tile) computeArea(other Tile) float64 {
	return (math.Abs(t.X-other.X) + 1) * (math.Abs(t.Y-other.Y) + 1)
}

type Node struct {
	Tile *Tile
	Next *Node
}

func NewNode(tile *Tile) *Node {
	return &Node{Tile: tile}
}

type rectBounds struct {
	xLo float64
	xHi float64
	yLo float64
	yHi float64
}

func newRectBounds(a, b Tile) rectBounds {
	xLo, xHi := a.X, b.X
	if xLo > xHi {
		xLo, xHi = xHi, xLo
	}

	yLo, yHi := a.Y, b.Y
	if yLo > yHi {
		yLo, yHi = yHi, yLo
	}

	return rectBounds{
		xLo: xLo,
		xHi: xHi,
		yLo: yLo,
		yHi: yHi,
	}
}

func (r rectBounds) contains(t *Tile) bool {
	return t.X > r.xLo && t.X < r.xHi && t.Y > r.yLo && t.Y < r.yHi
}

func (r rectBounds) crossesHorizontally(curr, next *Tile) bool {
	if curr.Y <= r.yLo || curr.Y >= r.yHi {
		return false
	}

	return (curr.X >= r.xHi && next.X <= r.xLo) || (curr.X <= r.xLo && next.X >= r.xHi)
}

func (r rectBounds) crossesVertically(curr, next *Tile) bool {
	if curr.X <= r.xLo || curr.X >= r.xHi {
		return false
	}

	return (curr.Y >= r.yHi && next.Y <= r.yLo) || (curr.Y <= r.yLo && next.Y >= r.yHi)
}

func (n *Node) formsValidRectangleWith(other *Node) bool {

	if n.Next == other {
		return true
	}

	bounds := newRectBounds(*n.Tile, *other.Tile)
	curr := n.Next

	for curr != n {
		if curr == other {
			curr = curr.Next
			continue
		}

		if bounds.contains(curr.Tile) {
			return false
		}

		if curr.Next == other {
			curr = curr.Next
			continue
		}

		if bounds.crossesHorizontally(curr.Tile, curr.Next.Tile) {
			return false
		}

		if bounds.crossesVertically(curr.Tile, curr.Next.Tile) {
			return false
		}

		curr = curr.Next
	}

	return true
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("usage: go run . <path/to/input/file> <constrain>")
	}

	constrain := os.Args[2] == "true"

	points, err := readPointsFromFile(os.Args[1])
	if err != nil {
		return err
	}

	pointPtrs := refutils.ToPointers(points)
	nodes := createNodeRing(pointPtrs)
	area := findLargestRectangle(nodes, constrain)

	fmt.Printf("The area of the largest rectange is %d\n", int(area))

	return nil
}

func findLargestRectangle(nodes []*Node, constrain bool) float64 {
	best := 0.0

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			delta := nodes[i].Tile.computeArea(*nodes[j].Tile)
			if delta > best {
				if !constrain || nodes[i].formsValidRectangleWith(nodes[j]) {
					best = delta
				}
			}
		}
	}
	return best
}

func createNodeRing(pointPtrs []*Tile) []*Node {

	nodes := make([]*Node, len(pointPtrs))
	for i, curr := range pointPtrs {
		nodes[i] = NewNode(curr)
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	nodes[len(nodes)-1].Next = nodes[0]

	return nodes
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
