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

func (dom DOM) Set(attribute string, value any) {

	dom.value.Set(attribute, value)
}

func New(value js.Value) DOM {

	return DOM{
		value: value,
	}
}
