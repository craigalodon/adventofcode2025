package main

import (
	"adventofcode2025/internal/mathutils"
	"adventofcode2025/internal/refutils"
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
		panic(fmt.Sprintf("invalid axis: %d", axis))
	}
}

func (t Tile) ForEachCoordinate(fn func(dimension int, value float64)) {
	fn(0, t.X)
	fn(1, t.Y)
}

func (t Tile) String() string {
	return fmt.Sprintf("(%f, %f)", t.X, t.Y)
}

func (t Tile) ComputeArea(other Tile) float64 {
	return (mathutils.Abs(t.X-other.X) + 1) * (mathutils.Abs(t.Y-other.Y) + 1)
}

func (t Tile) IsBetweenX(lo, hi Tile) bool {
	return t.X > lo.X && t.X < hi.X
}

func (t Tile) IsBetweenY(lo, hi Tile) bool {
	return t.Y > lo.Y && t.Y < hi.Y
}

func (t Tile) IsLeftOf(other Tile) bool {
	return t.X <= other.X
}

func (t Tile) IsRightOf(other Tile) bool {
	return t.X >= other.X
}

func (t Tile) IsBelow(other Tile) bool {
	return t.Y <= other.Y
}

func (t Tile) IsAbove(other Tile) bool {
	return t.Y >= other.Y
}

type Node struct {
	Tile *Tile
	Next *Node
}

func NewNode(tile *Tile) *Node {
	return &Node{Tile: tile}
}

func (n *Node) ValidRectangle(other *Node) bool {

	if n.Next == other {
		return true
	}

	right := other.Tile.IsRightOf(*n.Tile)
	above := other.Tile.IsAbove(*n.Tile)

	curr := n.Next

	for curr != n {
		if curr == other {
			curr = curr.Next
			continue
		}

		if right && above {
			if curr.Tile.IsBetweenX(*n.Tile, *other.Tile) && curr.Tile.IsBetweenY(*n.Tile, *other.Tile) {
				return false
			} else if curr.Next == other {
				goto Proceed
			} else if curr.Next.Tile.Y == curr.Tile.Y && curr.Tile.IsBetweenY(*n.Tile, *other.Tile) {
				if curr.Tile.IsRightOf(*other.Tile) && curr.Next.Tile.IsLeftOf(*n.Tile) {
					return false
				} else if curr.Tile.IsLeftOf(*n.Tile) && curr.Next.Tile.IsRightOf(*other.Tile) {
					return false
				}
			} else if curr.Tile.IsBetweenX(*n.Tile, *other.Tile) {
				if curr.Tile.IsAbove(*other.Tile) && curr.Next.Tile.IsBelow(*n.Tile) {
					return false
				} else if curr.Tile.IsBelow(*n.Tile) && curr.Next.Tile.IsAbove(*other.Tile) {
					return false
				}
			}
		} else if right {
			if curr.Tile.IsBetweenX(*n.Tile, *other.Tile) && curr.Tile.IsBetweenY(*other.Tile, *n.Tile) {
				return false
			} else if curr.Next == other {
				goto Proceed
			} else if curr.Next.Tile.Y == curr.Tile.Y && curr.Tile.IsBetweenY(*other.Tile, *n.Tile) {
				if curr.Tile.IsRightOf(*other.Tile) && curr.Next.Tile.IsLeftOf(*n.Tile) {
					return false
				} else if curr.Tile.IsLeftOf(*n.Tile) && curr.Next.Tile.IsRightOf(*other.Tile) {
					return false
				}
			} else if curr.Tile.IsBetweenX(*n.Tile, *other.Tile) {
				if curr.Tile.IsAbove(*n.Tile) && curr.Next.Tile.IsBelow(*other.Tile) {
					return false
				} else if curr.Tile.IsBelow(*other.Tile) && curr.Next.Tile.IsAbove(*n.Tile) {
					return false
				}
			}
		} else if above {
			if curr.Tile.IsBetweenX(*other.Tile, *n.Tile) && curr.Tile.IsBetweenY(*n.Tile, *other.Tile) {
				return false
			} else if curr.Next == other {
				goto Proceed
			} else if curr.Next.Tile.Y == curr.Tile.Y && curr.Tile.IsBetweenY(*n.Tile, *other.Tile) {
				if curr.Tile.IsRightOf(*n.Tile) && curr.Next.Tile.IsLeftOf(*other.Tile) {
					return false
				} else if curr.Tile.IsLeftOf(*other.Tile) && curr.Next.Tile.IsRightOf(*n.Tile) {
					return false
				}
			} else if curr.Tile.IsBetweenX(*other.Tile, *n.Tile) {
				if curr.Tile.IsAbove(*other.Tile) && curr.Next.Tile.IsBelow(*n.Tile) {
					return false
				} else if curr.Tile.IsBelow(*n.Tile) && curr.Next.Tile.IsAbove(*other.Tile) {
					return false
				}
			}
		} else {
			if curr.Tile.IsBetweenX(*other.Tile, *n.Tile) && curr.Tile.IsBetweenY(*other.Tile, *n.Tile) {
				return false
			} else if curr.Next == other {
				goto Proceed
			} else if curr.Next.Tile.Y == curr.Tile.Y && curr.Tile.IsBetweenY(*other.Tile, *n.Tile) {
				if curr.Tile.IsRightOf(*n.Tile) && curr.Next.Tile.IsLeftOf(*other.Tile) {
					return false
				} else if curr.Tile.IsLeftOf(*other.Tile) && curr.Next.Tile.IsRightOf(*n.Tile) {
					return false
				}
			} else if curr.Tile.IsBetweenX(*other.Tile, *n.Tile) {
				if curr.Tile.IsAbove(*n.Tile) && curr.Next.Tile.IsBelow(*other.Tile) {
					return false
				} else if curr.Tile.IsBelow(*other.Tile) && curr.Next.Tile.IsAbove(*n.Tile) {
					return false
				}
			}
		}

	Proceed:
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
	nodes := buildLinkedList(pointPtrs)
	area := findLargestRectangle(nodes, constrain)

	fmt.Printf("The area of the largest rectange is %d\n", int(area))

	return nil
}

func findLargestRectangle(nodes []*Node, constrain bool) float64 {
	best := 0.0

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			delta := nodes[i].Tile.ComputeArea(*nodes[j].Tile)
			if delta > best {
				if !constrain || nodes[i].ValidRectangle(nodes[j]) {
					best = delta
				}
			}
		}
	}
	return best
}

func buildLinkedList(pointPtrs []*Tile) []*Node {

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
