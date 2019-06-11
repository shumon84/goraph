package attribute

type Height struct {
	base
}

func NewHeight(f float64) *Height {
	height := &Height{}
	height.name = "height"
	height.AddValue(f)
	return height
}
