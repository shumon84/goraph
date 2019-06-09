package attribute

type Radius struct {
	base
}

func NewRadius(radius float64) *Radius {
	r := &Radius{}
	r.name = "r"
	r.AddValue(radius)
	return r
}
