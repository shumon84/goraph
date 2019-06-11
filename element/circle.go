package element

import "github.com/shumon84/goraph/attribute"

type Circle struct {
	base
}

func NewCircle(cx float64, cy float64, radius float64) *Circle {
	Circle := &Circle{}
	Circle.name = "circle"
	Circle.AddAttributes(attribute.NewCx(cx))
	Circle.AddAttributes(attribute.NewCy(cy))
	Circle.AddAttributes(attribute.NewRadius(radius))
	return Circle
}

func (Circle *Circle) String() string {
	return Circle.IndentString("")
}
