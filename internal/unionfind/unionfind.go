package unionfind

type Pair[T any] interface {
	First() T
	Second() T
}

func makeSet[T comparable](val T, parent map[T]T, rank map[T]int) {
	if _, exists := parent[val]; exists {
		return
	}
	parent[val] = val
	rank[val] = 0
}

func Find[T comparable](val T, parent map[T]T) T {
	if parent[val] != val {
		parent[val] = Find(parent[val], parent)
	}
	return parent[val]
}

func union[T comparable](v1, v2 T, parent map[T]T, rank map[T]int) {
	ra := Find(v1, parent)
	rb := Find(v2, parent)

	if ra == rb {
		return
	}

	rka := rank[ra]
	rkb := rank[rb]

	if rka < rkb {
		parent[ra] = rb
	} else if rka > rkb {
		parent[rb] = ra
	} else {
		parent[rb] = ra
		rank[ra] = rka + 1
	}
}

func BuildParentMap[T comparable](pairs []Pair[T]) map[T]T {
	parent := make(map[T]T)
	rank := make(map[T]int)

	for _, p := range pairs {
		makeSet(p.First(), parent, rank)
		makeSet(p.Second(), parent, rank)
		union(p.First(), p.Second(), parent, rank)
	}

	for k := range parent {
		Find(k, parent)
	}

	return parent
}
