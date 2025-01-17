package dom

import (
	"syscall/js"
)

type (
	DOM struct {
		value js.Value
	}
)

func (dom DOM) Get(attribute string) DOM {

	return New(dom.value.Get(attribute))
}

func New(value js.Value) DOM {

	return DOM{
		value: value,
	}
}
