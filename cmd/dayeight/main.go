package main

import (
	"adventofcode2025/internal/mathutils"
	"adventofcode2025/internal/spatial"
	"adventofcode2025/internal/unionfind"
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type JunctionBox struct {
	X float64
	Y float64
	Z float64
}

func (j JunctionBox) GetValue(depth int) float64 {

	axis := mathutils.Mod(depth, 3)

	switch axis {
	case 0:
		return j.X
	case 1:
		return j.Y
	default:
		return j.Z
	}
}

func (j JunctionBox) ForEachCoordinate(fn func(dimension int, value float64)) {
	fn(0, j.X)
	fn(1, j.Y)
	fn(2, j.Z)
}

type JunctionBoxPair struct {
	Left  JunctionBox
	Right JunctionBox
}

func (j JunctionBoxPair) First() JunctionBox {
	return j.Left
}

func (j JunctionBoxPair) Second() JunctionBox {
	return j.Right
}

func NewJunctionBoxPair(left, right *JunctionBox) JunctionBoxPair {
	return JunctionBoxPair{
		Left:  *left,
		Right: *right,
	}
}

type JunctionBoxPairDistance struct {
	Pair     JunctionBoxPair
	Distance float64
}

func NewJunctionBoxPairDistance(pair JunctionBoxPair, dist float64) JunctionBoxPairDistance {
	return JunctionBoxPairDistance{
		Pair:     pair,
		Distance: dist,
	}
}

type JunctionBoxPairDistMinHeap []JunctionBoxPairDistance

func (h *JunctionBoxPairDistMinHeap) Len() int {
	return len(*h)
}

func (h *JunctionBoxPairDistMinHeap) Less(i, j int) bool {
	return (*h)[i].Distance < (*h)[j].Distance
}

func (h *JunctionBoxPairDistMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *JunctionBoxPairDistMinHeap) Push(x any) {
	*h = append(*h, x.(JunctionBoxPairDistance))
}

func (h *JunctionBoxPairDistMinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type CircuitMaxHeap []int

func (h *CircuitMaxHeap) Len() int {
	return len(*h)
}

func (h *CircuitMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *CircuitMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *CircuitMaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *CircuitMaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func toPointers(boxes []JunctionBox) []*JunctionBox {
	result := make([]*JunctionBox, len(boxes))
	for i := range boxes {
		result[i] = &boxes[i]
	}
	return result
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 4 {
		return fmt.Errorf("usage: go run . <path/to/input/file> <k> <conn>")
	}

	points, err := readPointsFromFile(os.Args[1])
	if err != nil {
		return err
	}

	k, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return err
	}

	conn, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return err
	}

	pointPtrs := toPointers(points)
	root := spatial.KDTree(pointPtrs, 0)
	closest := getClosestPairs(pointPtrs, root, k)
	closestCopy := deepCopyValidHeap(closest)
	connected := getConnected(conn, closestCopy)
	parents := unionfind.BuildParentMap(connected)
	sorted := getSortedCircuitSizes(parents)
	vol := calculateCircuitVolume(sorted)

	fmt.Printf("The volume of the three largest circuits is %v\n", vol)

	dist, err := getDistanceFromWall(closest, points)
	if err != nil {
		return err
	}
	fmt.Printf("The distance from the wall is %d\n", dist)

	return nil
}

func getDistanceFromWall(closest *JunctionBoxPairDistMinHeap, points []JunctionBox) (int, error) {
	parent := make(map[JunctionBox]JunctionBox)
	rank := make(map[JunctionBox]int)
	components := len(points)

	for _, point := range points {
		parent[point] = point
		rank[point] = 0
	}

	for closest.Len() > 0 {
		curr := heap.Pop(closest).(JunctionBoxPairDistance)

		root1 := unionfind.Find(curr.Pair.Left, parent)
		root2 := unionfind.Find(curr.Pair.Right, parent)

		if root1 != root2 {
			if rank[root1] < rank[root2] {
				parent[root1] = root2
			} else if rank[root1] > rank[root2] {
				parent[root2] = root1
			} else {
				parent[root2] = root1
				rank[root1]++
			}
			components--

			if components == 1 {
				dist := int(curr.Pair.Left.X * curr.Pair.Right.X)
				return dist, nil
			}
		}
	}
	return 0, fmt.Errorf("unable to connect all junction boxes")
}

func deepCopyValidHeap(closest *JunctionBoxPairDistMinHeap) *JunctionBoxPairDistMinHeap {
	closestCopy := &JunctionBoxPairDistMinHeap{}
	*closestCopy = append(*closestCopy, *closest...)
	return closestCopy
}

func calculateCircuitVolume(sorted *CircuitMaxHeap) int {
	vol := 1
	for range 3 {
		if sorted.Len() > 0 {
			vol *= heap.Pop(sorted).(int)
		}
	}
	return vol
}

func getSortedCircuitSizes(parents map[JunctionBox]JunctionBox) *CircuitMaxHeap {
	counts := make(map[JunctionBox]int)

	for _, v := range parents {
		counts[v]++
	}

	h := &CircuitMaxHeap{}
	for _, v := range counts {
		heap.Push(h, v)
	}
	return h
}

func getConnected(conn int, closest *JunctionBoxPairDistMinHeap) []unionfind.Pair[JunctionBox] {
	connected := make([]unionfind.Pair[JunctionBox], 0, conn)
	for range conn {
		curr := heap.Pop(closest).(JunctionBoxPairDistance)
		connected = append(connected, curr.Pair)
	}
	return connected
}

func getClosestPairs(points []*JunctionBox, root *spatial.Node[JunctionBox], k int) *JunctionBoxPairDistMinHeap {
	h := &JunctionBoxPairDistMinHeap{}
	seen := make(map[JunctionBoxPair]bool)

	for _, p := range points {
		best := &spatial.NodeDistMaxHeap[JunctionBox]{}
		spatial.KNearestNeighbors(root, p, k, best, 0)
		for range k - 1 {
			curr := heap.Pop(best).(spatial.NodeDistance[JunctionBox])
			pair := NewJunctionBoxPair(p, curr.Node.Point)
			reversed := NewJunctionBoxPair(curr.Node.Point, p)
			if !seen[reversed] {
				heap.Push(h, NewJunctionBoxPairDistance(pair, curr.Distance))
				seen[pair] = true
			}
		}
	}

	return h
}

func readPointsFromFile(path string) ([]JunctionBox, error) {
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
	coords := make([]JunctionBox, 0)
	for scanner.Scan() {
		var c JunctionBox
		_, err := fmt.Sscanf(scanner.Text(), "%f,%f,%f", &c.X, &c.Y, &c.Z)
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
