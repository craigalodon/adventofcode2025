package mathutils

type Range struct {
	Lo int
	Hi int
}

func NewRange(lo, hi int) *Range {
	return &Range{lo, hi}
}
