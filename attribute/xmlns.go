package attribute

type XMLNS struct {
	base
}

func NewXMLNS(url string) *XMLNS {
	xmlns := &XMLNS{}
	xmlns.name = "xmlns"
	xmlns.AddValue(url)
	return xmlns
}
