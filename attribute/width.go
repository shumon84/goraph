package attribute

type Width struct {
	base
}

func NewWidth(f float64) *Width {
	width := &Width{}
	width.name = "width"
	width.AddValue(f)
	return width
}
