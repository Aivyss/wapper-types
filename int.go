package wrapperTypes

import (
	"errors"
	"fmt"
)

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

func (i *Int) Pow(num Int) Int {
	if num == 0 {
		return 1
	}

	if num < 0 {
		panic(errors.New("num < 0"))
	}

	result := *i
	for j := 0; j < num.Raw()-1; j++ {
		result *= *i
	}

	return result
}

func (i *Int) ToStr() String {
	return String(fmt.Sprintf("%v", i.Raw()))
}

func (i *Int) ToInt64() Int64 {
	return Int64(i.Raw())
}

func (i *Int) ToInt32() Int32 {
	return Int32(i.Raw())
}

func (i *Int) ToInt16() Int16 {
	return Int16(i.Raw())
}

func (i *Int) ToInt8() Int8 {
	return Int8(i.Raw())
}

func (i *Int) ToFloat64() Float64 {
	return Float64(i.Raw())
}

func (i *Int) ToFloat32() Float32 {
	return Float32(i.Raw())
}
