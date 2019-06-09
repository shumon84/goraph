package element

import (
	"github.com/shumon84/goraph/attribute"
)

type SVG struct {
	base
}

func NewSVG(width float64, height float64) *SVG {
	return NewSVGWithVersion(width, height, "http://www.w3.org/2000/svg")
}

func NewSVGWithVersion(width float64, height float64, url string) *SVG {
	svg := &SVG{}
	svg.name = "svg"
	svg.AddAttributes(attribute.NewWidth(width))
	svg.AddAttributes(attribute.NewHeight(height))
	svg.AddAttributes(attribute.NewXMLNS(url))
	return svg
}

func (svg *SVG) String() string {
	return svg.IndentString("")
}
