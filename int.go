package wrapperTypes

import "fmt"

type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64

func (i *Int) Raw() int {
	return int(*i)
}

func (i *Int) Abs() Int {
	raw := i.Raw()

	if raw < 0 {
		raw *= -1
	}

	return Int(raw)
}

func (i *Int) ToStr() String {
	return String(fmt.Sprintf("%v", i.Raw()))
}
