package mathutils

type Range struct {
	Lo int
	Hi int
}

func NewRange(lo, hi int) *Range {
	return &Range{lo, hi}
}

func (r *Range) Contains(v int) bool {
	return r.Lo <= v && v < r.Hi
}

func (r *Range) Size() int {
	return r.Hi - r.Lo
}

func (r *Range) TryUnionWith(o *Range) (*Range, bool) {
	if r.Lo > o.Hi || o.Lo > r.Hi {
		return nil, false
	}
	lo := min(r.Lo, o.Lo)
	hi := max(r.Hi, o.Hi)
	return NewRange(lo, hi), true
}

func GenerateCombinations(ranges []*Range) [][]int {
	var result [][]int
	var current []int

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(ranges) {
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}

		r := ranges[index]
		for i := r.Lo; i < r.Hi; i++ {
			current = append(current, i)
			backtrack(index + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}
