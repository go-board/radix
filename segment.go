package radix

type SegmentRadix struct {
	seg string
}

func NewSegment(s string) *SegmentRadix {
	return &SegmentRadix{seg: s}
}
