package attribute

type Cx struct {
	base
}

func NewCx(coordinateX float64) *Cx {
	cx := &Cx{}
	cx.name = "cx"
	cx.AddValue(coordinateX)
	return cx
}
