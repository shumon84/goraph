package element

import (
	"github.com/shumon84/goraph/attribute"
)

type Element interface {
	Name() string
	IndentString(string) string
	AddAttributes(...attribute.Attribute)
	AddElements(...Element)
}

type base struct {
	name       string
	attributes []attribute.Attribute
	elements   []Element
}

func (base *base) IndentString(indent string) string {
	str := indent + "<" + base.Name()
	for _, attr := range base.attributes {
		if attr == nil {
			continue
		}
		str += " " + attr.String()
	}

	if len(base.elements) == 0 {
		// if the number of element is 0, end this tag
		str += " />\n"
		return str
	}

	str += ">\n"

	for _, elm := range base.elements {
		if elm == nil {
			continue
		}
		str += elm.IndentString(indent + "  ")
	}

	str += indent + "</" + base.Name() + ">\n"
	return str
}

func (base *base) Name() string {
	return base.name
}

func (base *base) AddAttributes(attributes ...attribute.Attribute) {
	base.attributes = append(base.attributes, attributes...)
}

func (base *base) AddElements(elm ...Element) {
	base.elements = append(base.elements, elm...)
}
