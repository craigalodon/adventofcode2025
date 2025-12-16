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

func (r *Range) Union(o *Range) (*Range, bool) {
	if r.Lo > o.Hi || o.Lo > r.Hi {
		return nil, false
	}
	lo := min(r.Lo, o.Lo)
	hi := max(r.Hi, o.Hi)
	return NewRange(lo, hi), true
}
