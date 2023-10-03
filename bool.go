package wrapperTypes

import "fmt"

type Bool bool

func (b Bool) Raw() bool {
	return bool(b)
}

func (b Bool) ToStr() String {
	return String(fmt.Sprintf("%v", b.Raw()))
}

func (b Bool) ToInt() Int {
	if b {
		return 1
	}

	return 0
}

func (b Bool) And(c Bool) Bool {
	return b && c
}

func (b Bool) Or(c Bool) Bool {
	return b || c
}

func (b Bool) Not() Bool {
	return !b
}
