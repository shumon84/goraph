package attribute

import (
	"fmt"
)

type Attribute interface {
	fmt.Stringer
	Name() string
	AddValue(interface{})
}

type base struct {
	name  string
	value interface{}
}

func (base *base) String() string {
	return fmt.Sprint(base.Name(), `="`, base.value, `"`)
}

func (base *base) Name() string {
	return base.name
}

func (base *base) AddValue(value interface{}) {
	base.value = value
}
