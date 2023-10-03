package wrapperTypes

import "fmt"

type Bool bool

func (b Bool) Raw() bool {
	return bool(b)
}

func (b Bool) ToStr() String {
	return String(fmt.Sprintf("%v", b.Raw()))
}
