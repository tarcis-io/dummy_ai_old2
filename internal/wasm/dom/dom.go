package dom

import (
	"syscall/js"
)

type (
	DOM struct {
		value js.Value
	}
)

func New(value js.Value) DOM {

	return DOM{
		value: value,
	}
}
