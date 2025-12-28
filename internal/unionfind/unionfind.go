package unionfind

type Pair[T any] interface {
	First() T
	Second() T
}

func BuildParentMap[T comparable](pairs []Pair[T]) map[T]T {

	parent := make(map[T]T)
	rank := make(map[T]int)

	makeSet := func(val T) {
		if _, exists := parent[val]; exists {
			return
		}
		parent[val] = val
		rank[val] = 0
	}

	var find func(val T) T
	find = func(val T) T {
		if parent[val] != val {
			parent[val] = find(parent[val])
		}
		return parent[val]
	}

	union := func(v1, v2 T) {
		ra := find(v1)
		rb := find(v2)

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

	for _, p := range pairs {
		makeSet(p.First())
		makeSet(p.Second())
		union(p.First(), p.Second())
	}

	for k := range parent {
		find(k)
	}

	return parent
}
