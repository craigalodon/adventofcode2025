package spatial

import (
	"adventofcode2025/internal/mathutils"
	"container/heap"
	"math"
	"slices"
)

type Point interface {
	GetValue(depth int) float64
	ForEachCoordinate(fn func(dimension int, value float64))
}

func Distance[T Point](p1, p2 *T) float64 {
	res := 0.0

	(*p1).ForEachCoordinate(func(dim int, v1 float64) {
		v2 := (*p2).GetValue(dim)
		diff := v1 - v2
		res += diff * diff
	})

	return math.Sqrt(res)
}

type Node[T Point] struct {
	Point      *T
	LeftChild  *Node[T]
	RightChild *Node[T]
}

func KDTree[T Point](points []*T, depth int) *Node[T] {
	if len(points) == 0 {
		return nil
	}

	slices.SortFunc(points, func(a, b *T) int {
		v1 := (*a).GetValue(depth)
		v2 := (*b).GetValue(depth)
		return (int)(v1 - v2)
	})

	median := mathutils.FloorDiv(len(points), 2)

	node := Node[T]{
		Point:      points[median],
		LeftChild:  KDTree[T](points[:median], depth+1),
		RightChild: KDTree[T](points[median+1:], depth+1),
	}

	return &node
}

type NodeDistance[T Point] struct {
	Node     *Node[T]
	Distance float64
}

type NodeDistMaxHeap[T Point] []NodeDistance[T]

func (h *NodeDistMaxHeap[T]) Len() int {
	return len(*h)
}

func (h *NodeDistMaxHeap[T]) Less(i, j int) bool {
	return (*h)[i].Distance > (*h)[j].Distance
}

func (h *NodeDistMaxHeap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *NodeDistMaxHeap[T]) Push(x any) {
	*h = append(*h, x.(NodeDistance[T]))
}

func (h *NodeDistMaxHeap[T]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func KNearestNeighbors[T Point](root *Node[T], target *T, k int, h *NodeDistMaxHeap[T], depth int) {
	if root == nil {
		return
	}

	dist := Distance(root.Point, target)
	heap.Push(h, NodeDistance[T]{Node: root, Distance: dist})
	if h.Len() > k {
		heap.Pop(h)
	}

	v1 := (*root.Point).GetValue(depth)
	v2 := (*target).GetValue(depth)

	axisDist := v1 - v2

	var nextBranch, oppositeBranch *Node[T]
	if axisDist < 0 {
		nextBranch = root.LeftChild
		oppositeBranch = root.RightChild
	} else {
		nextBranch = root.RightChild
		oppositeBranch = root.LeftChild
	}

	KNearestNeighbors[T](nextBranch, target, k, h, depth+1)

	if axisDist < 0 {
		axisDist = -axisDist
	}

	if h.Len() < k || axisDist < (*h)[0].Distance {
		KNearestNeighbors[T](oppositeBranch, target, k, h, depth+1)
	}
}
