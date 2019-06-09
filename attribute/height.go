package attribute

type Height struct {
	base
}

func NewHeight(f float64) *Height {
	height := &Height{}
	height.name = "height"
	height.AddValue([]float64{f, f, f})
	return height
}
