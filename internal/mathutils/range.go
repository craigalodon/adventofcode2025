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
