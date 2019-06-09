package attribute

type Cy struct {
	base
}

func NewCy(coordinateY float64) *Cy {
	cy := &Cy{}
	cy.name = "cy"
	cy.AddValue(coordinateY)
	return cy
}
